package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Tree", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fileContent, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("error occured: ", error)
		os.Exit(1)
	}
	return strings.Split(string(fileContent), ",")
}

func (d deck) shuffle() deck {
	deckLength := len(d) - 1
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randomNumber := r.Intn(deckLength)
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
	return d
}
