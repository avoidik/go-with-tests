package main

import "testing"

func TestMemStore(t *testing.T) {

	t.Run("league", func(t *testing.T) {
		store := NewInMemStore()

		store.RecordWin("Jane")

		got := store.GetLeague()
		want := League{
			{Name: "Jane", Wins: 1},
		}
		assertLeague(t, got, want)
	})

	t.Run("score", func(t *testing.T) {
		store := NewInMemStore()

		store.RecordWin("Bill")
		got := store.GetPlayerScore("Bill")
		want := 1
		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}

		store.RecordWin("Bill")
		got = store.GetPlayerScore("Bill")
		want = 2
		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}
