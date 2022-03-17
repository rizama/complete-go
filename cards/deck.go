package main

import "fmt"

// Create new Type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of playing cards. Essentially an array of string
func newDeck() deck {
	// empty slice of deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards

}

// receiver (d deck)
func (d deck) myPrint() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Slice Range and Multiple Return
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
