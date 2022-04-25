package poker_test

import (
	"net/http"
	"net/http/httptest"
	"poker"
	"testing"
)

func TestPostGetMemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	fsps, err := poker.NewFileSystemPlayerStore(database)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	server := poker.NewPlayersServer(fsps)
	// server := NewPlayersServer(NewInMemStore())
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
		want := poker.League{
			{Name: "Fred", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}
