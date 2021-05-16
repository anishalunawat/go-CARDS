package main

func main() {
	cards := newDeckFromFile("a")
	cards.shuffle()
	/* Initialcards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	fmt.Println(cards.toString())
	cards.saveToFile("a")*/
	cards.print()

}
