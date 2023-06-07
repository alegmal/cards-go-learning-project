package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	
	if len(d) != 16 {
		t.Errorf("Expected deck lenghth of 20 but got %v", len(d))
	}

	expectedValue := "Ace of Spades"
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be %v but got %v", expectedValue, d[0])
	}

	expectedValue = "Four of Clubs"
	if d[len(d)-1] != expectedValue {
		t.Errorf("Expected first card to be %v but got %v", expectedValue, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.SaveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	expectedValue := 16
	if len(loadedDeck) != expectedValue {
		t.Errorf("Expected %d cards in deck, got %v", expectedValue, len(loadedDeck))
	}

	os.Remove(filename)
}