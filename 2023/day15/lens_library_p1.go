package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getAsciiValue(c rune) int {
	num, _ := strconv.Atoi(fmt.Sprintf("%d", c))
	return num
}

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	total := 0
	subTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		steps := strings.Split(line, ",")
		for _, step := range steps {
			subTotal = 0
			for _, char := range step {
				subTotal += getAsciiValue(char)
				subTotal *= 17
				subTotal = subTotal % 256
			}
			total += subTotal
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
