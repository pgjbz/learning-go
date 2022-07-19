package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //this thing tells to test suit "hey this pretty part of code is a helper".
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to peaple", func(t *testing.T) {
		got := Hello("Paulo", "English")
		want := "Hello, Paulo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in portuguese", func(t *testing.T) {
		got := Hello("Paulo", "Portuguese")
		want := "Olá, Paulo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Paulo", "Spanish")
		want := "Hola, Paulo"
		assertCorrectMessage(t, got, want)
	})

}
