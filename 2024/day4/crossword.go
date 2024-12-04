package main

import (
	"bufio"
	"fmt"
	"os"
)

type value struct {
	num string
	idx int
}

func solution1() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	dx := 0
	dy := 0
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		if dx == 0 {
			line := scanner.Text()
			dy = len(line)
		}
		dx += 1
	}

	puzzle := make([][]string, dy)
	for i := range puzzle {
		puzzle[i] = make([]string, dx)
	}

	inputs.Seek(0, 0)

	y := 0
	scanner = bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			puzzle[y][x] = string(c)
		}
		y += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	//words := make(map[string]int)

	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[0]); x++ {
			fmt.Println(x, y)
		}
	}

	return 0
}

func solution2() int {
	inputs, err := os.Open("sample.txt")
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
	total := solution1()
	fmt.Println("Part 1:", total)

	// part 2
	total = solution2()
	fmt.Println("Part 2:", total)
}
