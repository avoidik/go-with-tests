package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

const dbFilename = "game.db.json"

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}
	f.database.Encode(f.league)
}

func initializeDatabase(database *os.File) error {
	database.Seek(0, 0)

	dbInfo, err := database.Stat()
	if err != nil {
		return err
	}

	if dbInfo.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}

	return nil
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {

	err := initializeDatabase(database)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize database: %v", err)
	}

	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf("unable to decode database info: %v", err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&diskette{block: database}),
		league:   league,
	}, nil
}

func FileSystemPlayerStoreFromFile() (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open %q due to error %v", dbFilename, err)
	}
	closeFunc := func() {
		db.Close()
	}

	fsps, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("unexpected error: %v", err)
	}

	return fsps, closeFunc, nil
}
