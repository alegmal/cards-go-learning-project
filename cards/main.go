package main

import (
	"os"
)

func main() {
	filename := "test"

	cards := newDeck()
	cards.print()
	cards.SaveToFile(filename)
	newCards := newDeckFromFile(filename)
	newCards.shuffle()
	newCards.print()
	os.Remove(filename)
}
