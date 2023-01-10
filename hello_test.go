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
		got := Hello("world!", "English")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello you", func(t *testing.T) {
		got := Hello("you!", "English")
		want := "Hello, you!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish", func(t *testing.T) {
		got := Hello("Europa", "Spanish")
		want := "Hola, Europa"

		assertCorrectMessage(t, got, want)
	})
}
