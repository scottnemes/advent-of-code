package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	dish := [][]string{}
	roundRocks := map[string]bool{}

	scanner := bufio.NewScanner(inputs)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for x, char := range line {
			row = append(row, string(char))
			if string(char) == "O" {
				roundRocks[fmt.Sprintf("%d-%d", x, y)] = true
			}
		}
		dish = append(dish, row)
		y += 1
	}

	done := true
	for true {
		done = true
		for y, row := range dish {
			if y == 0 {
				continue
			}
			for x, char := range row {
				if char != "O" {
					continue
				}
				if dish[y-1][x] == "." {
					dish[y-1][x] = "O"
					dish[y][x] = "."
					done = false
				}
			}
		}
		if done {
			break
		}
	}

	total := 0
	mult := len(dish)
	for _, row := range dish {
		for _, char := range row {
			if char != "O" {
				continue
			}
			total += mult
		}
		mult -= 1
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
