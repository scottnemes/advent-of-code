package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution() (int) {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	total := 0

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
