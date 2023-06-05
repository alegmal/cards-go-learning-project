package main

import (
	"os"
)

// import (
// 	"fmt"
// )

func main() {
	filename := "test"
	cards := newDeck()
	cards.print()
	cards.SaveToFile(filename)
	newCards := newDeckFromFile("filenamef")
	newCards.print()
	os.Remove(filename)
}
