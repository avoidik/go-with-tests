package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostGet(t *testing.T) {
	server := NewPlayersServer(NewInMemStore())
	player := "Fred"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))

	t.Run("scores", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newGetScoreRequest(t, player))
		assertResponseCode(t, resp.Code, http.StatusOK)

		assertResponseBody(t, resp.Body.String(), "3")
	})

	t.Run("league", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newLeagueRequest(t))
		assertResponseCode(t, resp.Code, http.StatusOK)

		got := getLeagueFromResponse(t, resp.Body)
		want := []Player{
			{Name: "Fred", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}
