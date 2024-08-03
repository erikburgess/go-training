package main

import (
	"os"
	"testing"
)

// Tests should only test 1 thing
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("%s - %v", "Expected deck lenght of 52, but got", len(d))
	}

	// this should be a separate test
	if d[0] != "Spades Ace" {
		t.Errorf("%s - %v", " Expected first card of Ace of Spades, but got ", d[0])
	}

	// this should be a separate test
	if d[len(d)-1] != "Clubs King" {
		t.Errorf("%s - %v", "Expected last card to be of King of Clubs, but got", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFle(t *testing.T) {
	// TODO handle error
	_ = os.Remove("_decktesting")

	deck := newDeck()
	// TODO handle error
	_ = deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("%s - %v", " Execting 16 cards in deck, got", len(loadedDeck))
	}
}
