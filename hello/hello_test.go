package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, Word' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Joey", "French")
		want := "Bonjour, Joey"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// NOTE: t.Helper() is used to inform the test suite that this is a helper function. It ensures that the
	// line numbers in the output (if there is an error) come from the function call rather than our helper.
	// comment out helper and product an error and comment it back in to see the difference
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
