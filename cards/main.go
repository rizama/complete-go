package main

import "fmt"

func main() {
	cards := newDeck()

	//hand, remainingCards := deal(cards, 6)

	//hand.myPrint()
	//remainingCards.myPrint()

	//Save to file
	//errSaveFile := cards.saveToFile("cards.txt")
	//if errSaveFile != nil {
	//	return
	//}

	//Read from file
	//newCards := newDeckFromFile("cards.txt")
	//newCards.myPrint()

	//Shuffle
	cards.myPrint()
	fmt.Println("===============")
	cards.shuffle()
	cards.myPrint()
}
