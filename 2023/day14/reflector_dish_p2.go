package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rollDish(dir string, roundRocks map[string]bool, dish [][]string) {
	switch dir {
	case "north":
		rollDishN(roundRocks, dish)
	case "south":
		rollDishS(roundRocks, dish)
	case "west":
		rollDishW(roundRocks, dish)
	case "east":
		rollDishE(roundRocks, dish)
	}
}

func rollDishN(roundRocks map[string]bool, dish [][]string) {
	currentRow := 0
	for currentRow < len(dish) {
		for coords, active := range roundRocks {
			if !active {
				continue
			}
			x, _ := strconv.Atoi(strings.Split(coords, "-")[0])
			y, _ := strconv.Atoi(strings.Split(coords, "-")[1])
			if y != currentRow {
				continue
			}
			newY := y
			for row := y - 1; row >= 0; row -= 1 {
				if dish[row][x] != "." {
					break
				}
				newY = row
			}
			if newY == y {
				continue
			}
			// new rock pos
			roundRocks[fmt.Sprintf("%d-%d", x, newY)] = true
			dish[newY][x] = "O"
			// clear old rock pos
			roundRocks[coords] = false
			dish[y][x] = "."
		}
		currentRow += 1
	}
}

func rollDishS(roundRocks map[string]bool, dish [][]string) {
	currentRow := len(dish) - 1
	for currentRow >= 0 {
		for coords, active := range roundRocks {
			if !active {
				continue
			}
			x, _ := strconv.Atoi(strings.Split(coords, "-")[0])
			y, _ := strconv.Atoi(strings.Split(coords, "-")[1])
			if y != currentRow {
				continue
			}
			newY := y
			for row := y + 1; row < len(dish); row += 1 {
				if dish[row][x] != "." {
					break
				}
				newY = row
			}
			if newY == y {
				continue
			}
			// new rock pos
			roundRocks[fmt.Sprintf("%d-%d", x, newY)] = true
			dish[newY][x] = "O"
			// clear old rock pos
			roundRocks[coords] = false
			dish[y][x] = "."
		}
		currentRow -= 1
	}
}

func rollDishW(roundRocks map[string]bool, dish [][]string) {
	currentCol := 0
	for currentCol < len(dish[0]) {
		for coords, active := range roundRocks {
			if !active {
				continue
			}
			x, _ := strconv.Atoi(strings.Split(coords, "-")[0])
			y, _ := strconv.Atoi(strings.Split(coords, "-")[1])
			if x != currentCol {
				continue
			}
			newX := x
			for col := x - 1; col >= 0; col -= 1 {
				if dish[y][col] != "." {
					break
				}
				newX = col
			}
			if newX == x {
				continue
			}
			// new rock pos
			roundRocks[fmt.Sprintf("%d-%d", newX, y)] = true
			dish[y][newX] = "O"
			// clear old rock pos
			roundRocks[coords] = false
			dish[y][x] = "."
		}
		currentCol += 1
	}
}

func rollDishE(roundRocks map[string]bool, dish [][]string) {
	currentCol := len(dish[0]) - 1
	for currentCol >= 0 {
		for coords, active := range roundRocks {
			if !active {
				continue
			}
			x, _ := strconv.Atoi(strings.Split(coords, "-")[0])
			y, _ := strconv.Atoi(strings.Split(coords, "-")[1])
			if x != currentCol {
				continue
			}
			newX := x
			for col := x + 1; col < len(dish[0]); col += 1 {
				if dish[y][col] != "." {
					break
				}
				newX = col
			}
			if newX == x {
				continue
			}
			// new rock pos
			roundRocks[fmt.Sprintf("%d-%d", newX, y)] = true
			dish[y][newX] = "O"
			// clear old rock pos
			roundRocks[coords] = false
			dish[y][x] = "."
		}
		currentCol -= 1
	}
}

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

	directions := []string{"north", "west", "south", "east"}
	rounds := 1000000000
	for i := 0; i < rounds; i += 1 {
		for _, dir := range directions {
			rollDish(dir, roundRocks, dish)
		}
		fmt.Println("Done", i)
	}

	for _, row := range dish {
		fmt.Println(row)
	}

	total := 0
	// calculate total load
	for coords, active := range roundRocks {
		if !active {
			continue
		}
		y, _ := strconv.Atoi(strings.Split(coords, "-")[1])
		total += len(dish) - y
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
