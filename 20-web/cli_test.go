package poker_test

import (
	"fmt"
	"poker"
	"strings"
	"testing"
	"time"
)

type SpyBlindAlerter struct {
	alerts []poker.ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(alert poker.ScheduledAlert) {
	s.alerts = append(s.alerts, alert)
}

func TestCLI(t *testing.T) {

	t.Run("chris", func(t *testing.T) {
		store := &StubPlayerStore{
			scores:   map[string]int{},
			winClass: []string{},
		}

		in := strings.NewReader("Chris wins\n")

		spyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(store, in, spyAlerter)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Chris")
	})

	t.Run("cleo", func(t *testing.T) {
		store := &StubPlayerStore{
			scores:   map[string]int{},
			winClass: []string{},
		}

		in := strings.NewReader("Cleo wins\n")

		spyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(store, in, spyAlerter)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cleo")
	})

	t.Run("print blind values", func(t *testing.T) {
		store := &StubPlayerStore{
			scores:   map[string]int{},
			winClass: []string{},
		}

		in := strings.NewReader("Chris wins\n")
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(store, in, blindAlerter)
		cli.PlayPoker()

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Second, 200},
			{20 * time.Second, 300},
			{30 * time.Second, 400},
			{40 * time.Second, 500},
			{50 * time.Second, 600},
			{60 * time.Second, 800},
			{70 * time.Second, 1000},
			{80 * time.Second, 2000},
			{90 * time.Second, 4000},
			{100 * time.Second, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled at %v", c.Amount, c.At), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]
				assertScheduledAlert(t, alert, c)
			})
		}
	})
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got alert %v but want %v", got, want)
	}
}
