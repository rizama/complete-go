package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 16
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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove all document with this name
	os.Remove("_decktesting")

	// Generate new deck
	deck := newDeck()

	// Test Save to File
	deck.saveToFile("_decktesting")

	// Load Deck from File
	loadedDeck := newDeckFromFile("_decktesting")

	// Check Data
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	// Remove File
	os.Remove("_decktesting")
}
