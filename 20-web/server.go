package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(player string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) processWin(res http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	res.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(res http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		res.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(res, score)
}

func (p *PlayerServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	switch req.Method {
	case http.MethodGet:
		p.showScore(res, player)
	case http.MethodPost:
		p.processWin(res, player)
	}
}
