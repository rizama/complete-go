package main

import "fmt"

// var card string

func main() {
	// var card string = "Queens Heart"
	// card = "King Heart"

	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
