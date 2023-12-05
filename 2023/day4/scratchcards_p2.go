package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id int
	wins int
}

func solution() (int) {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	cards := []Card{}
	cardNum := 0
	winNums := []string{}
	ourNums := []string{}
	wins := 0

	allCards := []Card{}
	card := Card{}

	for scanner.Scan() {
		wins = 0
		line := scanner.Text()
		split := strings.Split(line, ": ")
		cardNum, _ = strconv.Atoi(strings.Fields(split[0])[1])
		winNums = strings.Split(strings.Split(split[1], " | ")[0], " ")
		ourNums = strings.Split(strings.Split(split[1], " | ")[1], " ")
		for _, num := range ourNums {
			if num != "" && slices.Contains(winNums, num) {
				wins += 1
			}
		}
		card = Card{}
		card.id = cardNum - 1
		card.wins = wins
		cards = append(cards, card)
	}

	p2Total := 0

	allCards = append(allCards, cards...)

	k := 0 // current card index
	for k < len(allCards) {
		card = allCards[k]
		wins = card.wins
		if wins > 0 {
			for i := 1; i <= wins; i += 1 {
				allCards = append(allCards, cards[card.id+i])
			}
		}
		k += 1
		p2Total += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return p2Total
}

func main() {
	p2Total := solution()
	fmt.Println("Part 2:", p2Total)
}
