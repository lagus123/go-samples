package main

func main() {
	cards := newDeck()
	hand, remaining := deal(cards, 4)

	hand.saveToFile("HandDeck")
	remaining.saveToFile("RemainingDeck")

	newCards := newDeckFromFile("RemainingDeck")
	newCards.print()

	cards.shuffle()
	cards.print()

}
