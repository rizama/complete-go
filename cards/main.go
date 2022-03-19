package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 6)

	hand.myPrint()
	remainingCards.myPrint()
	fmt.Println("==================================")
	fmt.Println(hand.toString())
}
