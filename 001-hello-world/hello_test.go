package main

import "testing"

func TestHello(t *testing.T) {

	assetsCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Hello with name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assetsCorrectMessage(t, got, want)
	})

	t.Run("Hello without name", func(t *testing.T) {
		got := Hello()
		want := "Hello, World"

		assetsCorrectMessage(t, got, want)
	})

	t.Run("Hello with many names use only first name", func(t *testing.T) {
		got := Hello("Chris", "John")
		want := "Hello, Chris"

		assetsCorrectMessage(t, got, want)
	})

}