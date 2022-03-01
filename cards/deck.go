package main

import "fmt"

// Create new Type of 'deck
// which is a slice of strings
type deck []string

func (d deck) myPrint() {
	for _, card := range d {
		fmt.Println(card)
	}
}
