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

func (d Deck) printDeck() {
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
			 4:"6",5:"7",6:"8",7:"9",
			 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}
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
	attr Attributes
}

func (h Hand) printHand() {
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
			 4:"6",5:"7",6:"8",7:"9",
			 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}

	fmt.Printf("%s: ", h.name)
	for _, card := range h.cards {
		fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
	}
	fmt.Println()
}

type Players []Hand

type Dealer Hand

type Cards []Card

type Combo []Cards


type Attributes struct {
	numCl			int	//number of cards in each suit
	numDi			int
	numHe			int
	numSp			int
	hasStrFlush		bool
	has4Kind		bool
	hasFullHouse	bool
	hasFlush		bool
	hasStraight		bool
	hasTrips		bool
	hasTwoPair		bool
	hasPair 		bool
	hasNoPair		bool
	handResult		string
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
	numPlayers := 5

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
	var dealer Hand
	dealer = dealGame(d, pl, dealer, numPlayers)
	evaluate(d, pl, dealer, numPlayers)
}
func evaluate(d Deck, pl Players, dealer Hand, numPlayers int) {
	var numSp, numHe, numDi, numCl int
	for card := range dealer.cards {
		switch dealer.cards[card].Suit {
		case 0: numCl++
		case 1: numDi++
		case 2: numHe++
		case 3: numSp++
		}
	}
	dealer.attr.numSp = numSp
	dealer.attr.numHe = numHe
	dealer.attr.numDi = numDi
	dealer.attr.numCl = numCl		
	fmt.Println(dealer.attr)
	for hand := range pl {
		numSp, numHe, numDi, numCl := 0, 0, 0, 0		
		for card := range pl[hand].cards {
			switch pl[hand].cards[card].Suit {
			case 0: numCl++
			case 1: numDi++
			case 2: numHe++
			case 3: numSp++
			}
		}
		pl[hand].attr.numSp = numSp
		pl[hand].attr.numHe = numHe
		pl[hand].attr.numDi = numDi
		pl[hand].attr.numCl = numCl
		//Determine hand value
		
		var possibleHands []Cards
		possibleHands = genHands(pl[hand], 2, dealer, 3)
		fmt.Println(possibleHands)
	}
}
/*
		select {
		case hasStrFlush(pl[hand], dealer):
			pl[hand].attr.handResult = "Straight Flush"
		case hasFullHouse(pl[hand], dealer):
			pl[hand].attr.handResult = "Full House"
		case hasFlush(pl[hand], dealer):
			pl[hand].attr.handResult = "Flush" 
		case hasStraight(pl[hand], dealer):
			pl[hand].attr.handResult = "Straight" 
		case hasTrips(pl[hand], dealer):
			pl[hand].attr.handResult = "Three of a kind" 
		case hasTwoPair(pl[hand], dealer): 
			pl[hand].attr.handResult = "Two Pair" 
		case hasPair(pl[hand], dealer):
			pl[hand].attr.handResult = "Pair" 
		case hasNoPair(pl[hand], dealer): 
			pl[hand].attr.handResult = "Nothing" }
		fmt.Println(pl[hand].attr)
	}
*/

/*
type Players []Hand

type Dealer Hand

type Cards []Card

type  []Cards []Cards
*/

func genHands(h1 Hand, numh1 int, h2 Hand, numh2 int )  []Cards {
	if (len(h1.cards) < numh1 || len(h2.cards) < numh2) {
		panic(fmt.Sprintf("%d %d Invalid integer parameter", numh1, numh2))
	}
	combo1 := combinations(h1, numh1)
	combo2 := combinations(h2, numh2)
	fmt.Println(combo1)
	fmt.Println(combo2)
	return combine(combo1, combo2)
}

func combine(combo1  []Cards, combo2  []Cards) []Cards { // to write, combine all possible combinations together
	var allCombos  []Cards
	return allCombos
}

func combinations(h Hand, r int)  []Cards {
	//Generates all combinations of hand h, choosing r (nCr)
	n := len(h.cards)	 
	var combos []Cards 	//total number of possible combinations of r (nCr) for 5C3
	hand := make(Cards, r)	//individual generated hands
	fmt.Printf("Entering combinationsUtil from combinations - hand= ")
	h.printHand()
	return combinationsUtil(h, combos, hand, 0, n-1, 0, r)
}

func combinationsUtil(h Hand, combos []Cards, hand Cards, start, end, index, r int) []Cards {

//	if start == 0 { indPos = 0 }
	if index == r { //hand has been generated
		combos = append(combos, hand)
		fmt.Println("Combos: ",combos)
		return combos
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
//		fmt.Printf("i= %d start= %d end= %d index= %d r= %d\n", i, start, end, index, r)
	   	hand[index] = h.cards[i]
    	combinationsUtil(h, combos, hand, i+1, end, index+1, r)
    }
    return combos
}


func dealGame(d Deck, pl Players, dealer Hand, numPlayers int) Hand {
	//deal dealer
	dealer.name = "Dealer"
	dealer.cards = d.cards[d.ind:d.ind+5]
	d.ind += 5
	//deal players
	for hand := range pl {
		pl[hand].cards = d.cards[d.ind:d.ind+4]
		d.ind += 4
	}
	displayGame(d, pl, dealer)
	return dealer
}

func displayGame(d Deck, pl Players, dealer Hand) {
	d.printDeck()
	dealer.printHand()
	//print Player's hands
	for _, player := range pl {
		player.printHand()
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

