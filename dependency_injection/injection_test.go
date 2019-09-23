package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jesse")

	got := buffer.String()
	want := "Hello, Jesse"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
