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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() deck {
	rSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rSource)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		shuffledCard := d[newPos]
		d[newPos] = d[i]
		d[i] = shuffledCard
	}
	return d
}
