package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type NormalSleep struct{}

type ConfigurableSleep struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (n *NormalSleep) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *ConfigurableSleep) Sleep() {
	s.sleep(s.duration)
}

func Countdown(w io.Writer, sleepFn Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleepFn.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}
	sleepFn.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleep{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
