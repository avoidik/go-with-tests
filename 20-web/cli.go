package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func (c *CLI) PlayPoker() {
	line := c.readLine()
	c.playerStore.RecordWin(strings.TrimSuffix(line, " wins"))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
