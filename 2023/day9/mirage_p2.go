package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	numbers := [][]int{}
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		l := []int{}
		n := strings.Split(line, " ")
		for _, j := range n {
			k, _ := strconv.Atoi(j)
			l = append(l, k)
		}
		numbers = append(numbers, l)
	}

	total := 0
	done := true
	for _, row := range numbers {
		diffs := [][]int{}
		diff := row
		for true {
			done = true
			r := []int{}
			for i := 1; i < len(diff); i += 1 {
				d := diff[i] - diff[i-1]
				r = append(r, d)
			}
			diffs = append(diffs, r)
			diff = diffs[len(diffs)-1]
			for _, n := range diff {
				if n != 0 {
					done = false
				}
			}
			if done {
				break
			}
		}
		firstNum := 0
		reversedDiffs := diffs
		slices.Reverse(reversedDiffs)
		for _, dRow := range reversedDiffs {
			firstNum = dRow[0] - firstNum
		}
		total += row[0] - firstNum
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
