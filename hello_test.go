package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello world", func(t *testing.T) {
		got := Hello("world!")
		want := "Hello, world!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Hello you", func(t *testing.T) {
		got := Hello("you!")
		want := "Hello, you!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
