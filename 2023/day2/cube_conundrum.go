package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func solution() (int, int) {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1, -1
	}
	defer inputs.Close()

	games := map[int]map[string]int{}
	game := map[string]int{}
	color := ""
	count := 0

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		game = map[string]int{
			"red": 0,
			"blue": 0,
			"green": 0,
		}
		line := scanner.Text()
		lineSplit := strings.Split(line, ": ")
		gameNum, _ := strconv.Atoi(strings.Split(lineSplit[0], " ")[1])
		plays := strings.Split(strings.Join(strings.Split(lineSplit[1], "; "), ", "), ", ")
		// stores the max seen amount of each cube color
		for _, v := range plays {
			color = strings.Split(v, " ")[1]
			count, _ = strconv.Atoi(strings.Split(v, " ")[0])
			game[color] = max(game[color], count)
		}
		games[gameNum] = game
	}

	target := map[string]int{
		"blue": 14,
		"green": 13,
		"red": 12,
	}

	p1Total := 0
	p2Total := 0
	match := true

	sub := 1

	// loop over the list of games and their cubes
	for g, p := range games {
		match = true
		sub = 1
		// loop over the required cubes
		// if we find a required cube that is not met, break early
		for c, n := range target {
			// multiply each max seen color cube count togther to get the power factor
			sub *= p[c]
			// if the max seen color cube count is greater-than the target, the game cannot be played
			if p[c] > n {
				match = false
			}
		}
		// if all of the required cubes are found, add the game number to the total
		if match {
			p1Total += g
		}
		p2Total += sub
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return p1Total, p2Total
}

func main() {
	// part 1
	p1Total, p2Total := solution()
	fmt.Println("Part 1:", p1Total)
	fmt.Println("Part 2:", p2Total)
}
