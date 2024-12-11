package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution1() int {
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
