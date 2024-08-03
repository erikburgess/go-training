package main

func main() {
	//cards := newDeck()

	//hand, remainingDeck := deal(cards, 5)
	//
	//fmt.Println(hand)
	//fmt.Println(remainingDeck)

	//fmt.Println(cards.toString())
	//err := cards.saveToFile("/tmp/cards.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	deck := newDeckFromFile("/tmp/cards.txt")
	deck.shuffle()
	deck.print()
	//fmt.Println(deck)

}
