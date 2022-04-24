package main

import "sync"

type InMemStore struct {
	score map[string]int
	mu    sync.RWMutex
}

func (s *InMemStore) GetPlayerScore(player string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.score[player]
}

func (s *InMemStore) RecordWin(player string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.score[player]++
}

func (s *InMemStore) GetLeague() League {
	var league League
	for name, wins := range s.score {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}

func NewInMemStore() *InMemStore {
	return &InMemStore{score: map[string]int{}}
}
