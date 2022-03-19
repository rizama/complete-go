package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 6)

	hand.myPrint()
	remainingCards.myPrint()

	err := cards.saveToFile("my_cardas.txt")
	fmt.Println(err)
	if err != nil {
		return
	}
}
