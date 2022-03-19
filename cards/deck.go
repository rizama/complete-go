package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// function with receiver (d deck)
func (d deck) myPrint() {
	if len(d) <= 0 {
		fmt.Println("No Deck to Print Out")
	} else {
		for _, card := range d {
			fmt.Println(card)
		}
	}
}

// Slice Range and Multiple Return (function with arguments)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Deck to String (function with receiver)
func (d deck) toString() string {
	// Joining a slice to string
	return strings.Join(d, ",")
}

//Function save file (function with receiver and arguments)
func (d deck) saveToFile(filename string) error {
	// filename, slice of bytes string, permission to create file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Function new DeckFromFile
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//	Option #1 -> Log the error and return an empty Deck
		fmt.Println("Error: ", err)
		return deck{}

		//	Option #2 -> Log the error and entirely quit the program
		//fmt.Println("Error: ", err)
		//os.Exit(1)
	}
	sliceOfString := strings.Split(string(bs), ",")
	return deck(sliceOfString)
}
