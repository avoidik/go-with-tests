package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "John", "Wins": 10},
			{"Name": "Jane", "Wins": 15}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := store.GetLeague()

		want := League{
			{Name: "Jane", Wins: 15},
			{Name: "John", Wins: 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "John", "Wins": 10},
			{"Name": "Jane", "Wins": 15}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := store.GetPlayerScore("Jane")
		want := 15

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}

		store.RecordWin("Jane")

		got = store.GetPlayerScore("Jane")
		want = 16

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("score new", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "John", "Wins": 10},
			{"Name": "Jane", "Wins": 15}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		store.RecordWin("Bill")

		got := store.GetPlayerScore("Bill")
		want := 1

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})

	t.Run("empty database", func(t *testing.T) {
		database, cleanUp := createTempFile(t, "")
		defer cleanUp()

		_, err := NewFileSystemPlayerStore(database)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
