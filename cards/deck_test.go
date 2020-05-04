package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first element card of Ace of Spades, but got %s", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected first element card of King of Clubes, but got %s", deck[0])
	}

}

func TestDeckFileMethods(t *testing.T) {
	os.Remove("__decktesting")

	deck := newDeck()
	deck.saveToFile("__decktesting")

	loadedDeck := newDeckFromFile("__decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards on deck")
	}

	os.Remove("__decktesting")
}
