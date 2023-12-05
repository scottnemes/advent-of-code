package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

	cards := map[int]int{}
	cardNum := 0
	winNums := []string{}
	ourNums := []string{}
	wins := 0

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
		// decrease cardNum by one to match 0-based indexing
		cardNum -= 1
		cards[cardNum] += 1
		for i := 1; i <= wins; i += 1 {
			cards[cardNum + i] += cards[cardNum]
		}
	}

	p2Total := 0
	for _, v := range cards {
		p2Total += v
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
