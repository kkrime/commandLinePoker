// Package hand provies a means to store and compare poker hands
package hand

import (
	"sort"
)

// cardValueMap stores a int value for each card
// with stronger cards having a higher numerical value
var cardValueMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

// results are stored as int values
type result int

const (
	Tie result = iota
	Win
	Lose
)

// handType represents the type of hand stored as a int value
// with strong hands having a higher numerical value
type handType int

const (
	highCard handType = iota
	pair
	twoPairs
	triple
	fullHouse
	fourOfAKind
)

type card struct {
	value    int
	quantity int
}

// Hand stores the cards and the type of hand
type Hand struct {
	handType handType
	cards    []card
}

// NewHand is used to create a new hand
// NewHand takes in a string of cards and returns
// a pointer to Hand
func NewHand(cards string) *Hand {

	hand := &Hand{}

	// map to count the number of occurences of each card
	// set size of map to 5 to minimize number of memory allocations to improve performance
	cardCountMap := make(map[rune]int, 5)

	for _, c := range cards {
		cardCountMap[c]++
	}

	// populate hand.card based on the values in cardCountMap
	hand.cards = make([]card, 0, len(cardCountMap))
	for c, quantity := range cardCountMap {
		hand.cards = append(hand.cards, card{value: cardValueMap[c], quantity: quantity})
	}

	// first sort by value in decending order
	sort.Slice(hand.cards, func(i, j int) bool {
		return hand.cards[i].value > hand.cards[j].value
	})

	// then sort by quantity in decending order
	sort.Slice(hand.cards, func(i, j int) bool {
		return hand.cards[i].quantity > hand.cards[j].quantity
	})

	// determine which type of card combination
	switch noOfDifferentCards := len(hand.cards); noOfDifferentCards {
	case 2:
		if hand.cards[0].quantity == 4 {
			hand.handType = fourOfAKind
		} else {
			hand.handType = fullHouse
		}
	case 3:
		if hand.cards[0].quantity == 3 {
			hand.handType = triple
		} else {
			hand.handType = twoPairs
		}
	case 4:
		hand.handType = pair
	case 5:
		hand.handType = highCard
	}

	return hand
}

// CompareHand compares the current hand to an opponents hand
// to detmine who wins, the result is relative to the reciver, not the opponnent
func (h *Hand) CompareHand(opponent *Hand) result {
	if h.handType > opponent.handType {
		return Win
	} else if h.handType < opponent.handType {
		return Lose
	} else {
		for i, _ := range h.cards {
			if h.cards[i].value > opponent.cards[i].value {
				return Win
			} else if h.cards[i].value < opponent.cards[i].value {
				return Lose
			}
		}
	}
	return Tie
}
