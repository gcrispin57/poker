package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"sort"
//	"reflect"
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

func print(i interface{}) { //prints type Hand, Deck, or C
	rankP := map[int] string{0:"2",1:"3",2:"4",3:"5",
			 4:"6",5:"7",6:"8",7:"9",
			 8:"T",9:"J",10:"Q",11:"K",12:"A"}
	suitP := map[int] string{0:"C",1:"D",2:"H",3:"S"}
//	var arr interface{}
	switch t := i.(type) {
	case Deck:
		var arr = Deck(t)
		fmt.Printf("%s", "Deck: ")
		for _, card := range arr.cards {
			fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
		}
		fmt.Println()
	case Hand:
		var arr = Hand(t)		
		fmt.Printf("%s: ", arr.name)
		for _, card := range arr.cards {
			fmt.Printf("%s%s ", rankP[card.Rank], suitP[card.Suit])
		}
		fmt.Println()
	case Combo:
		var arr = Combo(t)
		for _, cards := range arr {
			fmt.Printf(" [")
			for _, card := range cards {
				fmt.Printf("%s%s ",rankP[card.Rank], suitP[card.Suit])
			}
			fmt.Printf("]\n")
		}
		fmt.Println()
	case Cards:
		var arr = Cards(t)
		fmt.Printf("[")
		for i, card := range arr {
			if i != 0 && i != len(arr) {
				fmt.Printf(" ")
			}
			fmt.Printf("%s%s", rankP[card.Rank], suitP[card.Suit])
		}
		fmt.Printf("]")
	}
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
	handValue		int 	// numeric representation of handValue (ie Flush = 6, not evaluated = 0)
	handResult		string
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
}

// map of hand evaluations
func handValues() map[int]string {
	return map[int]string {
		0: "not evaluated",
		1: "No Pair",
		2: "Pair",
		3: "Two Pair",
		4: "Trips",
		5: "Straight",
		6: "Flush",
		7: "Full House",
		8: "4 of a Kind",
		9: "Straight Flush",
		10: "Royal Flush",
	} 
}

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	numPlayers := 5

//	deck := make([] Card, 52)
	var d Deck
	d.cards = make([]Card, 52)
	d.initDeck()
	print(d)
//	d.printDeck()
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
/*
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
*/

		//Determine hand value
	for hand := range pl {
		var possibleHands Combo
		possibleHands = genHands(pl[hand], 2, dealer, 3)
	//	possibleHands.printCombo()
		for hand := range possibleHands {
			var attr Attributes
			attr = getHandValue(possibleHands[hand], attr)
		}
	}
}

// Detemines the value of this specific hand and stores the result in attr
func getHandValue(hand Cards, attr Attributes) Attributes {
	// Map of Hand Values along with their String Representations

/*
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
*/


	var byRank = func(c1, c2 *Card) bool {
		if c1.Rank > c2.Rank {
			return true
		}
		if c1.Rank == c2.Rank && c1.Suit > c2.Suit {
			return true
		}
		return false
	}
//	var bySuit = func(c1, c2 *Card) bool {
//		return c1.Suit > c2.Suit
//	}
//	hand1 := Cards{Card{0,0}, Card{1,1}, Card{2,2}, Card{3,3}}
//	By(byRank).Sort(hand)
//		fmt.Println("By rank:", hand)
//	By(bySuit).Sort(hand)
//	fmt.Println("By suit:", hand)

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
//	fmt.Println(attr)

//	handVals := getHandValues()
// Flush and Straight and Straight Flush hand for testing= Cards{Card{0,0}, Card{1,0}, Card{2,0}, Card{3,0}, Card{4,0}}
// 4 of a kind: Cards{Card{9,0}, Card{9,1}, Card{9,2}, Card{9,3}, Card{4,0}}
// Full House: Cards{Card{9,0}, Card{9,1}, Card{4,2}, Card{4,3}, Card{4,0}}

//	fmt.Println("h ", reflect.TypeOf(hand), hand)
	print(hand)
	switch {
	case hasStrFlush(hand, byRank):
		fmt.Println("Straight Flush!!")
	case has4Kind(hand, byRank):
		fmt.Println("4 of a kind!!")
	case hasFullHouse(hand, byRank):
		fmt.Println("Full House!!")
	case hasFlush(hand):
		fmt.Println("Flush!!")
	case hasStraight(hand, byRank):
		fmt.Println("Straight!!")
	case hasTrips(hand, byRank):
		fmt.Println("Trips!!")
	case hasTwoPair(hand, byRank):
		fmt.Println("Two Pair!!")
	case hasPair(hand, byRank):
		fmt.Println("One Pair!!")
	case hasNoPair(hand, byRank):
		fmt.Println("No Pair!!")
	}
//	fmt.Println(hand)
	fmt.Println("next hand")
	return attr
}

func hasStrFlush(h Cards, byRank By) bool {
	return hasFlush(h) && hasStraight(h, byRank)
}

func has4Kind(h Cards, byRank By) bool {
	By(byRank).Sort(h)
	rank, cnt := 0, 0
	for i, card := range h {
		if i == 0 {
			rank = card.Rank
		} else {
			if card.Rank != rank {
				return false
			} else {
				cnt += 1
				if cnt == 3 { return true }
			}
		}
	}
	return false
}

func hasFullHouse(h Cards, byRank By) bool {
	By(byRank).Sort(h)
	if sameRank(h[:3], h[3:]) { return true }
	if sameRank(h[:2], h[2:]) { return true }
	return false
}

// verify if all supplied card slices have cards of the same Rank
func sameRank(arrCards ...Cards) bool {
	for _, cs := range arrCards {
		var rank = 0
		for i, c := range cs {
			if i == 0 {
				rank = c.Rank
			} else {
				if c.Rank != rank { return false }
			}
		}
	}
	return true
}

func hasStraight(h Cards, byRank By) bool {
	By(byRank).Sort(h)
	rank := 0
	for i, card := range h {
		if i == 0 {
			rank = card.Rank
		} else {
			if card.Rank != rank - 1 { return false }
			rank -= 1
		}
	}
	return true
}

func hasFlush(h Cards) bool {
//	var flushSuit int
	flushSuit := 0
	for i, card := range h {
		if i == 0 {
			flushSuit = card.Suit
		} else {
			if card.Suit != flushSuit {
				return false
			}
		}
	}
	return true
}

func hasTrips(h Cards, byRank By) bool {
	By(byRank).Sort(h)
	if has4Kind(h, byRank) || hasFullHouse(h, byRank) { return false }
	if sameRank(h[:3]) || sameRank(h[2:]) { return true }
	return false
}

func hasTwoPair(h Cards, byRank By) bool {
	By(byRank).Sort(h)
	if has4Kind(h, byRank) || hasFullHouse(h, byRank) ||
		hasTrips(h, byRank) { return false }
	cnt := 0
	for i := 0; i + 2 <= len(h); i++ {
		if sameRank(h[i:i+2]) { 
			cnt += 1
			i += 2
		 }		
	}
	if cnt == 2 { return true }
	return false
}

func hasPair(h Cards, byRank By) bool {
	By(byRank).Sort(h)
	if has4Kind(h, byRank) || hasFullHouse(h, byRank) ||
		hasTrips(h, byRank) || hasTwoPair(h, byRank) { return false }
	cnt := 0
	for i := 0; i + 2 <= len(h); i++ {
		if sameRank(h[i:i+2]) { 
			cnt += 1
			i += 2
		 }		
	}
	if cnt == 1 { return true }
	return false
}

func hasNoPair(h Cards, byRank By) bool {
	if hasStrFlush(h, byRank) || has4Kind(h, byRank) ||
		hasFullHouse(h, byRank) || hasFlush(h) ||
		hasStraight(h, byRank) || hasTrips(h, byRank) ||
		 hasTwoPair(h, byRank) {
		  return false
	}
	return true
}

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
//	print(allCombos)
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
		print(h)
//		h.printHand()
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
//			h.printHand()
			print(h)
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
			print(h)
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
	print(d)
	print(dealer)
	//print Player's hands
	for _, player := range pl {
		print(player)
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

