package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostGet(t *testing.T) {
	store := NewInMemStore()
	server := PlayerServer{store: store}
	player := "Fred"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(t, player))

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, newGetScoreRequest(t, player))
	assertResponseCode(t, resp.Code, http.StatusOK)

	assertResponseBody(t, resp.Body.String(), "3")
}
