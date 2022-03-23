package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 20
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "Four of Clubs"

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected deck length of %v, but got %v", expectedFirstCard, d[0])
	}

	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected deck length of %v, but got %v", expectedLastCard, d[len(d)-1])
	}
}
