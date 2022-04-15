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

func NewInMemStore() *InMemStore {
	return &InMemStore{score: map[string]int{}}
}

func main() {
	store := &PlayerServer{store: NewInMemStore()}
	log.Fatal(http.ListenAndServe(":5000", store))
}
