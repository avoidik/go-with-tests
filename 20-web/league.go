package poker

import (
	"encoding/json"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(r io.Reader) (League, error) {
	var league League
	if err := json.NewDecoder(r).Decode(&league); err != nil {
		return nil, err
	}
	return league, nil
}
