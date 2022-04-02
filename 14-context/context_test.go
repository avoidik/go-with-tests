package contextsrv

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWritter struct {
	written bool
}

func (s *SpyResponseWritter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWritter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWritter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	data := "Hello, world!"

	t.Run("basic test", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}

		srv := Server(store)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)

		if resp.Body.String() != data {
			t.Errorf("expected %q got %q", data, resp.Body.String())
		}
	})
	t.Run("cancel earlier", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}

		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelCtx, cancelFunc := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancelFunc)
		req = req.WithContext(cancelCtx)

		resp := &SpyResponseWritter{}
		srv.ServeHTTP(resp, req)

		if resp.written {
			t.Error("no response expected for cancelled request")
		}
	})
}
