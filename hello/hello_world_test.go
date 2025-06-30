package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got, want := Hello("Chris", ""), "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello', World' when empty string is supplied", func(t *testing.T) {
		got, want := Hello("", ""), "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got, want := Hello("Elodie", "Spanish"), "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got, want := Hello("Ellis", "French"), "Bonjour, Ellis"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}