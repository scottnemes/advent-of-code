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
	number int
	winning []string
	current []string
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
	cardNum := ""
	winNums := []string{}
	ourNums := []string{}
	
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		cardNum = strings.Fields(split[0])[1]
		winNums = strings.Split(strings.Split(split[1], " | ")[0], " ")
		ourNums = strings.Split(strings.Split(split[1], " | ")[1], " ")
		card := Card{}
		card.winning = winNums
		card.current = ourNums
		card.number, _ = strconv.Atoi(cardNum)
		card.number -= 1
		cards = append(cards, card)
	}

	p2Total := 0

	allCards := []Card{}
	allCards = append(allCards, cards...)

	wins := 0
	k := 0 // current card index
	for k < len(allCards) {
		card := allCards[k]
		for _, num := range card.current {
			if num != "" && slices.Contains(card.winning, num) {
				wins += 1
			}
		}
		if wins > 0 {
			for i := 1; i <= wins; i += 1 {
				allCards = append(allCards, cards[card.number + i])
			}
		}
		p2Total += 1
		wins = 0
		k += 1
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
