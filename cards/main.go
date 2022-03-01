package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Heart")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards.myPrint()
}

func newCard() string {
	return "Five of Diamonds"
}
