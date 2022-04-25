package poker_test

import (
	"poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("chris", func(t *testing.T) {
		store := &StubPlayerStore{
			scores:   map[string]int{},
			winClass: []string{},
		}

		in := strings.NewReader("Chris wins\n")

		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Chris")
	})

	t.Run("cleo", func(t *testing.T) {
		store := &StubPlayerStore{
			scores:   map[string]int{},
			winClass: []string{},
		}

		in := strings.NewReader("Cleo wins\n")

		cli := poker.NewCLI(store, in)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cleo")
	})
}
