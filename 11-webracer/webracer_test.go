package webracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(td time.Duration) *httptest.Server {
	delayedServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(td)
		rw.WriteHeader(http.StatusOK)
	}))
	return delayedServer
}

func TestWebsiteRacer(t *testing.T) {
	t.Run("both servers delaying", func(t *testing.T) {
		slowServer := makeDelayedServer(70 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(50 * time.Millisecond)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		got, err := WebsiteRacer(slowUrl, fastUrl)
		want := slowUrl
		if err != nil {
			t.Errorf("no error expected, but returned")
		}
		if got == want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
	t.Run("error if longer 10 sec", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		slowUrl := slowServer.URL
		fastUrl := slowServer.URL

		_, err := ConfigurableWebsiteRacer(slowUrl, fastUrl, 10*time.Millisecond)
		if err == nil {
			t.Errorf("expected timeout error, but none thrown")
		}
	})
}
