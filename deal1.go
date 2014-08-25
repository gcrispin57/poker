package main

import (
	"fmt"
	"math/rand"
	"time"
)


type Card struct {
	Rank int
	Suit int
}


type Deck struct {
	deck []Card
	ind int // points to "top" of deck
}

func (d Deck) initDeck() {
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			d.deck[j*13 + i] = Card{i, j}
		}
	}
	d.shuffle()
}

func (d Deck) printDeck() {
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
							 4:"6",5:"7",6:"8",7:"9",
							 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}
	for _, card := range d.deck {
		fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
	}
}

func (d Deck) shuffle() {
	var rnd int
	for i := range d.deck {
		rnd = rand.Intn(51)
		d.deck[i], d.deck[rnd] = d.deck[rnd], d.deck[i]
	}
}


/*
func New(deck interface{}) (Deck error) {
		deck = make([] Card, 52)
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			deck[j*13 + i] = Card{i, j}
		}
	}
}


func (deck Deck) NewDeck() {
	deck = {Card{0,0}, Card{0,1}}
//	return deck
}
*/

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
//	deck := make([] Card, 52)
	var d Deck
	d.deck = make([]Card, 52)
	d.initDeck()
	d.printDeck()
}


//	deck.Deal()
//	deck.Print()

