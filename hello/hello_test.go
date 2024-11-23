package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("test Bisha", func(t *testing.T) {
		want := "Hello, Bisha"
		got := Hello("Bisha", "")
		checkMessage(t, want, got)

	})

	t.Run("test Dasha", func(t *testing.T) {
		want := "Hello, Dasha"
		got := Hello("Dasha", "")
		checkMessage(t, want, got)
	})

	t.Run("test empty", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")
		checkMessage(t, want, got)
	})

	t.Run("in Spanish", func(t *testing.T) {
		want := "Hola, Grisha"
		got := Hello("Grisha", "Spanish")
		checkMessage(t, want, got)
	})

	t.Run("in German", func(t *testing.T) {
		want := "Heil, Grunken"
		got := Hello("Grunken", "German")
		checkMessage(t, want, got)
	})

}

func checkMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q, but got %q", want, got)
	}

}
