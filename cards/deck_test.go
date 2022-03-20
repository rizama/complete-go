package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 16

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}
}
