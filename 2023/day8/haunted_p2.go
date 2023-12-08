package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func solution() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	turns := []string{}
	nodes := map[string][]string{}
	startNodes := map[string][]string{}
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
		if string(split[0][2]) == "A" {
			startNodes[split[0]] = append(startNodes[split[0]], strings.TrimPrefix(split2[0], "("))
			startNodes[split[0]] = append(startNodes[split[0]], strings.TrimSuffix(split2[1], ")"))
		}
	}

	total := 0
	turnIdx := 0
	nextNodes := map[string]string{}
	nextNode := ""
	done := true
	matchCount := 0
	maxCount := 0
	var wg sync.WaitGroup
	// keep looping until we find the end
	for true {
		done = true
		matchCount = 0
		for node, _ := range startNodes {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if nextNodes[node] == "" {
					nextNode = nodes[node][directions[turns[turnIdx]]]
				} else {
					nextNode = nodes[nextNodes[node]][directions[turns[turnIdx]]]
				}
				nextNodes[node] = nextNode
				if !strings.HasSuffix(nextNode, "Z") {
					done = false
				} else {
					matchCount += 1
				}
			}()
			wg.Wait()
		}
		if matchCount > maxCount {
			maxCount = matchCount
			fmt.Println(maxCount, nextNodes)
		}
		turnIdx += 1
		// if we've exhausted the list of turns, start over
		if turnIdx == len(turns) {
			turnIdx = 0
		}
		total += 1
		// check to see if all paths have landed on an ending node
		if done {
			break
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
