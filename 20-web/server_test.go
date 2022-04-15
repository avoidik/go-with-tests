package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winClass []string
	league   []Player
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
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostScoreRequest(t *testing.T, player string) *http.Request {
	t.Helper()
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winClass = append(s.winClass, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestPostPlayers(t *testing.T) {
	storeStub := &StubPlayerStore{
		map[string]int{},
		[]string{},
		[]Player{},
	}

	server := NewPlayersServer(storeStub)

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
	server := NewPlayersServer(&StubPlayerStore{
		map[string]int{
			"Floyd":  10,
			"Pepper": 20,
		},
		[]string{},
		[]Player{},
	})

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

func newLeagueRequest(t *testing.T) *http.Request {
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func assertContentJson(t *testing.T, headers *http.Header) {
	t.Helper()
	want := "application/json"
	got := headers.Get("content-type")
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func getLeagueFromResponse(t *testing.T, body *bytes.Buffer) []Player {
	t.Helper()
	var got []Player
	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("unable to decode %q with err: %v", body, err)
	}
	return got
}

func TestGetLeague(t *testing.T) {
	t.Run("get players", func(t *testing.T) {
		wantedLeague := []Player{
			{Name: "Peter", Wins: 10},
			{Name: "Derek", Wins: 15},
			{Name: "Mary", Wins: 7},
		}
		store := &StubPlayerStore{}
		store.league = wantedLeague
		server := NewPlayersServer(store)

		req := newLeagueRequest(t)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertContentJson(t, &res.Result().Header)

		got := getLeagueFromResponse(t, res.Body)
		assertResponseCode(t, res.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}
