package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winClass []string
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func newGetScoreRequest(t *testing.T, player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostScoreRequest(t *testing.T, player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winClass = append(s.winClass, name)
}

func TestPostPlayers(t *testing.T) {
	storeStub := &StubPlayerStore{
		map[string]int{},
		[]string{},
	}

	server := &PlayerServer{store: storeStub}

	t.Run("Mary", func(t *testing.T) {
		req := newPostScoreRequest(t, "Mary")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseCode(t, res.Code, http.StatusAccepted)

		if len(storeStub.winClass) != 1 {
			t.Errorf("got %d but want %d calls", len(storeStub.winClass), 1)
		}

		if storeStub.winClass[0] != "Mary" {
			t.Errorf("got %q but expected %q", storeStub.winClass[0], "Mary")
		}
	})
}

func TestGetPlayers(t *testing.T) {
	server := &PlayerServer{store: &StubPlayerStore{
		map[string]int{
			"Floyd":  10,
			"Pepper": 20,
		},
		[]string{},
	}}

	t.Run("404", func(t *testing.T) {
		req := newGetScoreRequest(t, "Fred")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseCode(t, res.Code, http.StatusNotFound)
	})

	t.Run("Pepper", func(t *testing.T) {
		req := newGetScoreRequest(t, "Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseCode(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "20")
	})

	t.Run("Floyd", func(t *testing.T) {
		req := newGetScoreRequest(t, "Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseCode(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "10")
	})
}
