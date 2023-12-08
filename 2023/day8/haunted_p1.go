package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	turns := []string{}
	nodes := map[string][]string{}
	directions := map[string]int{
		"L": 0,
		"R": 1,
	}

	x := 0
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		x += 1
		line := scanner.Text()
		// skip empty lines
		if line == "" {
			continue
		}
		// store turning directions
		if x == 1 {
			for _, dir := range line {
				turns = append(turns, string(dir))
			}
			continue
		}
		// store nodes
		split := strings.Split(line, " = ")
		split2 := strings.Split(split[1], ", ")
		nodes[split[0]] = append(nodes[split[0]], strings.TrimPrefix(split2[0], "("))
		nodes[split[0]] = append(nodes[split[0]], strings.TrimSuffix(split2[1], ")"))
	}

	total := 0
	currentNode := "AAA"
	turnIdx := 0
	// keep looping until we find the end
	for currentNode != "ZZZ" {
		currentNode = nodes[currentNode][directions[turns[turnIdx]]]
		turnIdx += 1
		// if we've exhausted the list of turns, start over
		if turnIdx == len(turns) {
			turnIdx = 0
		}
		total += 1
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
