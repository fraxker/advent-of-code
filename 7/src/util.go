package src

import (
	"fmt"
	"log"
)

type Camel struct {
	Hand *Hand
	Bid int
}

func (c *Camel) String() string {
	return fmt.Sprintf("Bid: %d, Hand: %v", c.Bid, c.Hand)
}

type Hand struct {
	Cards []int
	Count map[int]int
}

func NewHand() *Hand {
	return &Hand{Cards: make([]int, 0, 5), Count: make(map[int]int)}
}

func (h *Hand) Add(card int) {
	if len(h.Cards) == 5 {
		log.Fatal("Hand already full")
	}
	h.Cards = append(h.Cards, card)
	h.Count[card]++
}

func (h *Hand) AddJokers() {
	// Get Jokers from map
	jokers := h.Count[1]
	delete(h.Count, 1)

	// Get high count key from map
	var highCount int
	var highKey int
	for key, count := range h.Count {
		if count > highCount || (count == highCount && key > highKey){
			highCount = count
			highKey = key
		}
	}

	// Add Jokers to high count key
	h.Count[highKey] += jokers
}

func (h *Hand) Rank() int {
	h.AddJokers()
	// Five of a kind
	if len(h.Count) == 1 {
		return 7
	} else if len(h.Count) == 2 {
		// Four of a kind
		for _, count := range h.Count {
			if count == 4 {
				return 6
			}
		}
		// Full house
		return 5
	} else if len(h.Count) == 3 {
		// Three of a kind
		for _, count := range h.Count {
			if count == 3 {
				return 4
			}
		}
		// Two pair
		return 3
	} else if len(h.Count) == 4 {
		//One pair
		return 2
	} 

	// High card
	return 1
}

func (h *Hand) String() string {
	return fmt.Sprintf("Cards: %v, Count: %v", h.Cards, h.Count)
}

type ByCamel []*Camel

func (b ByCamel) Len() int {
	return len(b)
}	

func (b ByCamel) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}


func (b ByCamel) Less(i, j int) bool {
	// Compare Hand Ranks
	if b[i].Hand.Rank() != b[j].Hand.Rank() {
		return b[i].Hand.Rank() < b[j].Hand.Rank()
	}
	// If equal, compare card order
	for idx := 0; idx < 5; idx++ {
		if b[i].Hand.Cards[idx] != b[j].Hand.Cards[idx] {
			return b[i].Hand.Cards[idx] < b[j].Hand.Cards[idx]
		}
	}
	return false
}