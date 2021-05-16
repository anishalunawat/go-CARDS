package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is slice of type string

type deck []string

// (d deck) is a way of defining  receiver here to
// tell that print method can be applied to this type
// go is not an object oriented language and
// "this a way we can bind methods to a type"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spade", "Heart", "Diamond", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

func deal(d deck, hands int) (deck, deck) {
	return d[:hands], d[hands:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := (ioutil.ReadFile(fileName))
	if err != nil {
		// Option #1 - log error and return call to newDeck()
		// Option #2 - log error and quit the program
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return s
}

func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPositions := r.Intn(len(d) - 1)
		d[i], d[newPositions] = d[newPositions], d[i]
	}
}
