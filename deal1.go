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


type Deck []Card

func (deck Deck) initDeck() {
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			deck[j*13 + i] = Card{i, j}
		}
	}
	deck.shuffle()
}

func (deck Deck) printDeck() {
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
							 4:"6",5:"7",6:"8",7:"9",
							 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}
	for _, card := range deck {
		fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
	}
}

func (deck Deck) shuffle() {
	var rnd int
	for i := range deck {
		rnd = rand.Intn(51)
		deck[i], deck[rnd] = deck[rnd], deck[i]
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
	var deck Deck
	deck = make([]Card, 52)
	deck.initDeck()
	deck.printDeck()
}


//	deck.Deal()
//	deck.Print()

