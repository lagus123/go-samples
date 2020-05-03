package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // Return new slice

	for _, card := range cards {
		fmt.Println(card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
