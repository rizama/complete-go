package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Heart")

	cards.myPrint()
}

func newCard() string {
	return "Five of Diamonds"
}
