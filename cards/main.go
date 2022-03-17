package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 6)

	hand.myPrint()
	remainingCards.myPrint()
}
