package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}

	}

	return cards
}

func (d deck) saveToFile(filename string) {
	temp := toByteSlice(d)
	err := ioutil.WriteFile(filename, temp, 0666)
	check(err)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	check(err)

	return strings.Split(string(bs), ",")

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPos := r.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]

	}

}

func toByteSlice(d deck) []byte {
	return []byte(strings.Join(d, ","))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
