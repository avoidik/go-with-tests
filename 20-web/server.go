package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func (p *PlayerServer) processWin(res http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprint(res, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) showScore(res http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		res.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(res, score)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := p.store.GetLeague()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leagueTable)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func NewPlayersServer(store PlayerStore) *PlayerServer {
	server := &PlayerServer{}
	server.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(server.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(server.playersHandler))
	server.Handler = router
	return server
}
