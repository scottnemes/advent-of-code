package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type StartNode struct {
	lastCount   int
	lastTurnIdx int
	nextNode    string
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	turns := []string{}
	nodes := map[string][]string{}
	startNodes := map[string]StartNode{}
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
			startNode := StartNode{}
			startNode.lastCount = 0
			startNode.lastTurnIdx = 0
			startNode.nextNode = ""
			startNodes[split[0]] = startNode
		}
	}

	nextNode := ""
	counts := []int{}
	var wg sync.WaitGroup
	for node, _ := range startNodes {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			sn := startNodes[n]
			i := 0
			j := 0
			nextNode = sn.nextNode
			for true {
				if nextNode == "" {
					nextNode = nodes[n][directions[turns[j]]]
				} else {
					nextNode = nodes[sn.nextNode][directions[turns[j]]]
				}
				i += 1
				j += 1
				// if we've exhausted the list of turns, start over
				if j >= len(turns) {
					j = 0
				}
				sn.nextNode = nextNode
				if strings.HasSuffix(nextNode, "Z") {
					sn.lastCount = i
					sn.lastTurnIdx = j
					startNodes[node] = sn
					counts = append(counts, sn.lastCount)
					return
				}
			}
		}(node)
		wg.Wait()
	}

	countLcm := lcm(counts[0], counts[1])
	for i := 2; i < len(counts); i += 1 {
		countLcm = lcm(countLcm, counts[i])
	}

	return countLcm
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
