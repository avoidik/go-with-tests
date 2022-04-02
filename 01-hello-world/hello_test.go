package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to a person", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world if empty parameter is given", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", langSpanish)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Claude", langFrench)
		want := "Bonjour, Claude"

		assertCorrectMessage(t, got, want)
	})
}
