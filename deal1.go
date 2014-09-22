package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"sort"
)

type Card struct {
	Rank int
	Suit int
}


// By is the type of a "less" function that defines the ordering of its Card arguments.
type By func(c1, c2 *Card) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(cards Cards) {
	cs := &cardSorter{
		cards: cards,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(cs)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type cardSorter struct {
	cards Cards
	by      func(p1, p2 *Card) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *cardSorter) Len() int {
	return len(s.cards)
}

// Swap is part of sort.Interface.
func (s *cardSorter) Swap(i, j int) {
	s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *cardSorter) Less(i, j int) bool {
	return s.by(&s.cards[i], &s.cards[j])
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

func (c Combo) printCombo() {
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
			 4:"6",5:"7",6:"8",7:"9",
			 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}

	for _, cards := range c {
		fmt.Printf(" [")
		for _, card := range cards {
			fmt.Printf("%s%s ",rankP[card.Rank], suitP[card.Suit])
		}
		fmt.Printf("]\n")
	}
	fmt.Println()
}


func (cards Cards) Len() int {
	return len(cards)
}

func (cards Cards) Less(i, j int) bool {
    if cards[i].Suit > cards[j].Suit {
        return true 
    } else {
        if cards[i].Suit == cards[j].Suit && cards[i].Rank > cards[j].Rank {
            return true
        } 
    }
    return false
}

func (cards Cards) Swap(i, j int) {
    cards[i], cards[j] = cards[j], cards[i]
}

func (c Cards) String() string {
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
		 4:"6",5:"7",6:"8",7:"9",
		 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}
	str := "["
	for i, card := range c {
		if i > 0 {
			str += " "
		}
		str += rankP[card.Rank]
		str += suitP[card.Suit]
	}
	str += "]"
	return str
}

type Players []Hand

type Dealer Hand

type Cards []Card

type Combo []Cards

type HandEval struct {
	
}

type Attributes struct {
	numSp			int	//number of cards in each suit
	numHe			int
	numDi			int
	numCl			int
	hasRoyalFlush	bool
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
		var possibleHands Combo
		possibleHands = genHands(pl[hand], 2, dealer, 3)
	//	possibleHands.printCombo()
		for hand := range possibleHands {
			var attr Attributes
			attr = getAttributes(possibleHands[hand], attr)
		}
	}
}


func getAttributes(hand Cards, attr Attributes) Attributes {
	var byRank = func(c1, c2 *Card) bool {
		if c1.Rank > c2.Rank {
			return true
		}
		if c1.Rank == c2.Rank && c1.Suit > c2.Suit {
			return true
		}
		return false
	}
	var bySuit = func(c1, c2 *Card) bool {
		return c1.Suit > c2.Suit
	}
//	hand1 := Cards{Card{0,0}, Card{1,1}, Card{2,2}, Card{3,3}}
	fmt.Println("hand:", hand)
	By(byRank).Sort(hand)
		fmt.Println("By rank:", hand)
	By(bySuit).Sort(hand)
	fmt.Println("By suit:", hand)

	var numSp, numHe, numDi, numCl int
	for card := range hand {
		switch hand[card].Suit {
		case 0: numCl++
		case 1: numDi++
		case 2: numHe++
		case 3: numSp++
		}
	}
	attr.numSp = numSp
	attr.numHe = numHe
	attr.numDi = numDi
	attr.numCl = numCl
	fmt.Println(attr)
//	fmt.Println(hand)
	fmt.Println("next hand")
	return attr
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
//	fmt.Println(combo1)
	combo2 := combinations(h2, numh2)
//	fmt.Println(combo2)
	allCombos := combine(combo1, combo2)
//	fmt.Println(allCombos)
	return allCombos
}

func combine(combo1  []Cards, combo2  []Cards) []Cards { // to write, combine all possible combinations together
	allCombos := make([]Cards, 60)
	ind := 0
	for c1 := range combo1 {
		for c2 := range combo2 {
			combo3 := combo1[c1]
			for cardInd := range combo2[c2] {
				combo3 = append(combo3, combo2[c2][cardInd])
			}
			allCombos[ind] = combo3
			ind += 1
//			fmt.Println(combo1[c1], combo1[c1][0])
//			fmt.Println(combo2[c2])
//
		}
	}
	return allCombos
}

func combinations(h Hand, r int)  []Cards {
	//Generates all combinations of hand h, choosing r (nCr)
	n := len(h.cards)
	var combos []Cards
	var cards Cards
	if r == 3 && n == 5 { // Dealer combinations 5C3 10 combos in total
		h.printHand()
		combos = make([]Cards, 10)
		ind := 0
		for i := 0; i < 3; i++ {
			for j := i+1; j < 4; j++ {
				for k := j+1; k < 5; k++ {
					cards = make(Cards, 3 )
					cards[0] = h.cards[i] 
					cards[1] = h.cards[j]
					cards[2] = h.cards[k]
					combos[ind] = cards
					ind = ind + 1
				}
			}
		} 
	} else {
		if r == 2 && n == 4 { // player hand 4C2 6 combos
			h.printHand()
			combos = make([]Cards, 6)
			var ind int = 0
			for i := 0; i < 3; i++ {
				for j := i+1; j < 4; j++ {
					cards = make(Cards, 2)
					cards[0] = h.cards[i] 
					cards[1] = h.cards[j]
					combos[ind] = cards
					ind += 1
				}
			}
		} else {
			hand := make(Cards, r)	//individual generated hands
			fmt.Printf("Entering combinationsUtil from combinations - hand= ")
			h.printHand()
			c := make(chan []Card)
			go combinationsUtil(h, combos, hand, 0, n-1, 0, r, c)
			for h1 := range c {
        		println( "recv from c", h1 )
				combos = append(combos, h1)        		
    		}
		}
	}
	return combos
}

func combinationsUtil(h Hand, combos []Cards, hand Cards, start, end, index, r int, c chan []Card) {

//	if start == 0 { indPos = 0 }
	if index == r { //hand has been generated
		fmt.Println(hand)
		c <- hand
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
//		fmt.Printf("i= %d start= %d end= %d index= %d r= %d\n", i, start, end, index, r)
	   	hand[index] = h.cards[i]
	   	combinationsUtil(h, combos, hand, i+1, end, index+1, r, c)
    }
    close(c)
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

