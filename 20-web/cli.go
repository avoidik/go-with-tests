package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		alerter:     alerter,
	}
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(ScheduledAlert{At: blindTime, Amount: blind})
		blindTime = blindTime + (10 * time.Second)
	}
}

func (c *CLI) PlayPoker() {
	c.scheduleBlindAlerts()
	line := c.readLine()
	c.playerStore.RecordWin(strings.TrimSuffix(line, " wins"))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
