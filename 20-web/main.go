package main

import (
	"log"
	"net/http"
	"sync"
)

type InMemStore struct {
	score map[string]int
	mu    sync.Mutex
}

func (s *InMemStore) GetPlayerScore(player string) int {
	return s.score[player]
}

func (s *InMemStore) RecordWin(player string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.score[player]++
}

func (s *InMemStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.score {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}

func NewInMemStore() *InMemStore {
	return &InMemStore{score: map[string]int{}}
}

func main() {
	server := NewPlayersServer(NewInMemStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
