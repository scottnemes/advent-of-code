package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func solution() (int) {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	cards := map[string]map[string][]string{}
	cardNum := ""
	winNums := []string{}
	ourNums := []string{}
	
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		cardNum = split[0]
		winNums = strings.Split(strings.Split(split[1], " | ")[0], " ")
		ourNums = strings.Split(strings.Split(split[1], " | ")[1], " ")
		card := map[string][]string{}
		card["winning"] = winNums
		card["numbers"] = ourNums
		cards[cardNum] = card
	}

	p1Total := 0
	p1CardTotal := 0

	for _, card := range cards {
		for _, num := range card["numbers"] {
			if num != "" && slices.Contains(card["winning"], num) {
				p1CardTotal = max(1, p1CardTotal * 2)
			}
		}
		p1Total += p1CardTotal
		p1CardTotal = 0
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return p1Total
}

func main() {
	p1Total := solution()
	fmt.Println("Part 1:", p1Total)
}
