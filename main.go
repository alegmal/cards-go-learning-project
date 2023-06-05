package main

// import (
// 	"fmt"
// )

func main() {
	cards := newDeck()
	cards.print()
	cards.SaveToFile("test")
	newCards := newDeckFromFile("test")
	newCards.print()
}
