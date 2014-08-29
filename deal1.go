package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type Card struct {
	Rank int
	Suit int
}


type Deck struct {
	cards []Card
	ind int // points to "top" of deck
}

func (d Deck) initDeck() {
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			d.cards[j*13 + i] = Card{i, j}
		}
	}
	d.shuffle()
}

func (d Deck) printDeck(rankP, suitP map[int]string) {
	fmt.Printf("%s", "Deck: ")
	for _, card := range d.cards {
		fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
	}
	fmt.Println()
}

func (d Deck) shuffle() {
	var rnd int
	for i := range d.cards {
		rnd = rand.Intn(51)
		d.cards[i], d.cards[rnd] = d.cards[rnd], d.cards[i]
	}
}

type Hand struct {
	cards []Card
	name string
}

func (h Hand) printHand(rankP, suitP map[int]string) {
	fmt.Printf("%s: ", h.name)
	for _, card := range h.cards {
		fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
	}
	fmt.Println()
}

type Players []Hand

type Dealer Hand

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
	numPlayers := 5

	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
			 4:"6",5:"7",6:"8",7:"9",
			 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}

//	deck := make([] Card, 52)
	var d Deck
	d.cards = make([]Card, 52)
	d.initDeck()
//	d.printDeck(rankP, suitP)
	var pl Players
	pl = make([]Hand, numPlayers)
	var h Hand
	for i:=0; i < numPlayers; i++ {
		h.cards = make([]Card, 4)
		h.name = "player " + strconv.Itoa(i)
		pl[i] = h
	}
	dealGame(d, pl, numPlayers, rankP, suitP)
}


func dealGame(d Deck, pl Players, numPlayers int, rankP, suitP map[int]string) {
	//deal dealer
	var dealer Hand
	dealer.name = "Dealer"
	dealer.cards = d.cards[d.ind:d.ind+5]
	d.ind += 5
	//deal players
	for hand := range pl {
		pl[hand].cards = d.cards[d.ind:d.ind+4]
		d.ind += 4
	}
	displayGame(d, pl, dealer, rankP, suitP)
}

func displayGame(d Deck, pl Players, dealer Hand, rankP, suitP map[int]string) {
	d.printDeck(rankP, suitP)
	dealer.printHand(rankP, suitP)
	//print Player's hands
	for _, player := range pl {
		player.printHand(rankP, suitP)
	}
}

/*
func createPlayers(pl Players, n int) {
	var hand Hand
	for i := 0; i < n; i++ {
		hand.cards := make([]Card, 4)
		hand.name := "player" + string(i)
		pl[i] := hand
	}
}
*/


//	deck.Deal()
//	deck.Print()

