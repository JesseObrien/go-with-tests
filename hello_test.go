package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bob")
	want := "Hello, Bob"

	if got != want {
		t.Errorf("go %q want %q", got, want)
	}
}
