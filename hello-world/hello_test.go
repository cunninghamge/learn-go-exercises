package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gus")
	want := "Hello, Gus"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
