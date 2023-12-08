package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2

// Five of a kind, where all five cards have the same label: AAAAA
// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// High card, where all cards' labels are distinct: 23456

type Card struct {
	name string
	value int
}

type Hand struct {
	cards []Card
	bid int
	counts map[string]int
	strongest int // best hand that can be made from these cards
}

type Game struct {
	hands []Hand
}

var types = map[string]int {
	"five-of-a-kind": 6,
	"four-of-a-kind": 5,
	"full-house": 4,
	"three-of-a-kind": 3,
	"two-pair": 2,
	"one-pair": 1,
	"high-card": 0,
}

func (h *Hand) getStrongestHand() {
	strongest := -1
	twoKind := 0
	threeKind := 0
	for card, count := range h.counts {
		// don't process jokers like normal cards
		if card == "J" {
			continue
		}
		switch count {
		case 5:
			strongest = types["five-of-a-kind"]
		case 4:
			strongest = max(strongest, types["four-of-a-kind"])
		case 3:
			threeKind += 1
			strongest = max(strongest, types["three-of-a-kind"])
		case 2:
			twoKind += 1
			strongest = max(strongest, types["one-pair"])
		}
	}

	// pair-related matchings
	if threeKind == 1 && twoKind == 1 {
		// full house
		strongest = max(strongest, types["full-house"])
	} else if twoKind == 2 {
		// two pair
		strongest = max(strongest, types["two-pair"])
	} else if strongest == -1 {
		strongest = types["high-card"]
	}

	// adjust for the joker, it's wild
	for i := 0; i < h.counts["J"]; i += 1 {
		switch strongest {
		case types["five-of-a-kind"]:
			// already the max, stop processing jokers
			break
		case types["four-of-a-kind"]:
			strongest = types["five-of-a-kind"]
		case types["full-house"]:
			strongest = types["four-of-a-kind"]
		case types["three-of-a-kind"]:
			strongest = types["four-of-a-kind"]
		case types["two-pair"]:
			strongest = types["full-house"]
		case types["one-pair"]:
			strongest = types["three-of-a-kind"]
		case types["high-card"]:
			strongest = types["one-pair"]
		}
	}
	h.strongest = strongest
}

func (g *Game) sort() bool {
	changed := false
	sort.Slice(g.hands, func(i, j int) bool { return g.hands[i].strongest < g.hands[j].strongest })
	sorted := []Hand{}
	for i, hand := range g.hands {
		// if this is the first hand, add it to start
		if len(sorted) == 0 {
			sorted = append(sorted, hand)
			continue
		}
		// if this hand is stronger than the previous hand, add it
		if hand.strongest > sorted[i-1].strongest {
			sorted = append(sorted, hand)
			continue
		}
		// similar strength hands, deep dive
		for j, _ := range hand.cards {
			if hand.cards[j].value < sorted[i-1].cards[j].value {
				// card of this hand is greater-than card of last sorted hand, swap them
				sorted = append(sorted[:len(sorted)-1], hand, sorted[i-1])
				changed = true
				break
			} else if hand.cards[j].value > sorted[i-1].cards[j].value {
				// card of this hand is less-than card of last sorted hand, add this one after
				sorted = append(sorted, hand)
				break
			} else {
				// card of this hand is the same as the card of the last sorted hand, check next card
				continue
			}
		}
	}
	g.hands = sorted
	return changed
}

func solution() (int) {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	cardValues := map[string]int {
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 1,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	game := Game{}
	hand := Hand{}
	card := Card{}
	split := []string{}
	bid := 0
	cards := ""
	for scanner.Scan() {
		line := scanner.Text()
		hand = Hand{}
		hand.counts = map[string]int{}
		split = strings.Split(line, " ")
		cards = string(split[0])
		bid, _ = strconv.Atoi(string(split[1]))
		hand.bid = bid
		for _, c := range cards {
			card = Card{}
			card.name = string(c)
			card.value = cardValues[card.name]
			hand.counts[card.name] += 1
			hand.cards = append(hand.cards, card)
		}
		hand.getStrongestHand()
		game.hands = append(game.hands, hand)
	}

	// keep sorting the list until there are no changes
	for game.sort() { }

	total := 0

	for i, h := range game.hands {
		total += h.bid * (i + 1)
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
