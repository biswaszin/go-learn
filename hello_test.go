package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


// This is a testing file. And it's really important.
