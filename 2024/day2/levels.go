package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type value struct {
	num string
	idx int
}

// Safe:
// The levels are either all increasing or all decreasing
// Any two adjacent levels differ by at least one and at most three

func solution1() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		increasing := false
		decreasing := false
		unsafe := false
		for i := range levels {
			if i == 0 {
				continue
			}
			cur, _ := strconv.Atoi(levels[i])
			prev, _ := strconv.Atoi(levels[i-1])
			diff := cur - prev
			if diff == 0 || math.Abs(float64(diff)) > 3 {
				unsafe = true
				break
			}
			if increasing && diff < 0 {
				unsafe = true
				break
			}
			if decreasing && diff > 0 {
				unsafe = true
				break
			}
			if !increasing && !decreasing {
				if diff < 0 {
					decreasing = true
				} else {
					increasing = true
				}
				continue
			}
		}
		if !unsafe {
			total += 1
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return total
}

func solution2() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		original_levels := append([]string(nil), levels...)
		for j := 0; j <= len(original_levels); j++ {
			increasing := false
			decreasing := false
			unsafe := false
			for i := range levels {
				if i == 0 {
					continue
				}
				cur, _ := strconv.Atoi(levels[i])
				prev, _ := strconv.Atoi(levels[i-1])
				diff := cur - prev
				if i == 1 {
					if diff < 0 {
						decreasing = true
					} else {
						increasing = true
					}
				}
				if diff == 0 || math.Abs(float64(diff)) > 3 {
					unsafe = true
				} else if increasing && diff < 0 {
					unsafe = true
				} else if decreasing && diff > 0 {
					unsafe = true
				}
				if unsafe {
					break
				}
			}
			if !unsafe {
				total += 1
				break
			}
			if j == len(original_levels) {
				break
			}
			levels = append([]string(nil), original_levels...)
			levels = slices.Delete(levels, j, j+1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return total
}

func main() {
	// part 1
	total := solution1()
	fmt.Println("Part 1:", total)

	// part 2
	total = solution2()
	fmt.Println("Part 2:", total)
}
