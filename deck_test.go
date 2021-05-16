package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length was 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected First card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected Last card of Four of Club but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards im deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
