package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calcDistance(a []int, b []int) float64 {
	return math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1]))
}

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}

	defer inputs.Close()
	galixies := [][]string{}
	i := 1
	y := 0
	empty := true
	expansionSize := 1000000 - 1

	emptyRows := []int{}
	emptyCols := []int{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		empty = true
		line := scanner.Text()
		row := []string{}
		for _, char := range line {
			newChar := ""
			if string(char) != "." {
				empty = false
				newChar = strconv.Itoa(i)
				i += 1
			} else {
				newChar = string(char)
			}
			row = append(row, string(newChar))
		}
		galixies = append(galixies, row)
		if empty {
			emptyRows = append(emptyRows, y)
		}
		y += 1
	}

	for x := 0; x < len(galixies[0]); x += 1 {
		empty := true
		for y := 0; y < len(galixies); y += 1 {
			if galixies[y][x] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}

	positions := map[int][]int{}
	for x := 0; x < len(galixies[0]); x += 1 {
		for y := 0; y < len(galixies); y += 1 {
			if galixies[y][x] != "." {
				pos, _ := strconv.Atoi(galixies[y][x])
				positions[pos] = []int{x, y}
			}
		}
	}

	expandPositions := map[int][]int{}
	for k, v := range positions {
		expandPositions[k] = append(expandPositions[k], v...)
	}

	// expand galaxies
	for _, y := range emptyRows {
		for k, coords := range positions {
			if coords[1] > y {
				expandPositions[k][1] += expansionSize
			}
		}
	}
	for _, x := range emptyCols {
		for k, coords := range positions {
			if coords[0] > x {
				expandPositions[k][0] += expansionSize
			}
		}
	}

	distances := map[string]int{}
	total := 0

	for num1, coords1 := range expandPositions {
		for num2, coords2 := range expandPositions {
			if num1 == num2 {
				continue
			}
			str1 := fmt.Sprintf("%d-%d", num1, num2)
			str2 := fmt.Sprintf("%d-%d", num2, num1)
			// check if this pair has been calced already
			_, exists1 := distances[str1]
			_, exists2 := distances[str2]
			if exists1 || exists2 {
				continue
			}
			dist := int(calcDistance(coords1, coords2))
			distances[str1] = dist
			total += dist
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
