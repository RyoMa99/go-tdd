package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Hello world", func(t *testing.T) {
		got := Hello("world!")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello you", func(t *testing.T) {
		got := Hello("you!")
		want := "Hello, you!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
