package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const opSleep = "sleep"
const opWrite = "write"

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, opSleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, opWrite)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("basic output", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleep := &SpyCountdownOperations{}
		Countdown(&buffer, spySleep)
		got := buffer.String()
		want := "3\n2\n1\nGo!"
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}
		if len(spySleep.Calls) != 4 {
			t.Errorf("expected %d got %d", 4, len(spySleep.Calls))
		}
	})
	t.Run("specific order output", func(t *testing.T) {
		spySleep := &SpyCountdownOperations{}
		Countdown(spySleep, spySleep)
		got := spySleep.Calls
		want := []string{
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %q got %q", want, got)
		}
	})
	t.Run("configurable sleep", func(t *testing.T) {
		spyTime := &SpyTime{}
		spySleep := ConfigurableSleep{5 * time.Second, spyTime.Sleep}
		spySleep.Sleep()
		if spyTime.durationSlept != spySleep.duration {
			t.Errorf("expected to sleep %v but got %v", spySleep.duration, spyTime.durationSlept)
		}
	})
}
