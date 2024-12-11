package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.

// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
// The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone.
// (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

func solution1() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	pebbles := []int{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			pebbles = append(pebbles, n)
		}
	}

	blinks := 75
	for i := 0; i < blinks; i += 1 {
		fmt.Println("Blink", i)
		idx := 0
		pebbles_copy := make([]int, len(pebbles))
		copy(pebbles_copy, pebbles)
		for _, pebble := range pebbles_copy {
			if pebble == 0 {
				pebbles[idx] = 1
				idx += 1
				continue
			}
			pebble_str := strconv.Itoa(pebble)
			if len(pebble_str)%2 == 0 {
				l := len(pebble_str) / 2
				left, _ := strconv.Atoi(pebble_str[:l])
				right, _ := strconv.Atoi(pebble_str[l:])
				tail := make([]int, len(pebbles[idx+1:]))
				copy(tail, pebbles[idx+1:])
				pebbles = append(pebbles[:idx], left, right)
				pebbles = append(pebbles, tail...)
				idx += 2
				continue
			}
			pebbles[idx] = pebbles[idx] * 2024
			idx += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return len(pebbles)
}

func solution2() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return 0
}

func main() {
	// part 1
	start := time.Now()
	total := solution1()
	fmt.Println("Part 1:", total, time.Since(start))

	// part 2
	start = time.Now()
	total = solution2()
	fmt.Println("Part 2:", total, time.Since(start))
}
