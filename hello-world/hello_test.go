package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World tested function"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
