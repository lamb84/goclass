package main

import (
	"fmt"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuite := []string{"diamond", "spade", "heart", "club"}
	cardNumber := []string{"ace", "two", "three", "four"}

	for _, suite := range cardSuite {
		for _, number := range cardNumber {
			cards = append(cards, number+" of "+suite)
		}
	}
	return cards
}

func newHand(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
