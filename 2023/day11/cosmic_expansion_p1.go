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
	//return math.Sqrt(math.Abs(float64((b[0]-a[0])^2)) + math.Abs(float64((b[1]-a[1])^2)))
}

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}

	defer inputs.Close()
	scanner := bufio.NewScanner(inputs)
	galixies := [][]string{}
	i := 0
	y := 0
	empty := true
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
			galixies = append(galixies, row)
			y += 1
		}
		y += 1
	}

	emptyCols := []int{}

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

	for y, _ := range galixies {
		loopCount := 0
		for _, x := range emptyCols {
			galixies[y] = append(galixies[y][:x+1+loopCount], galixies[y][x+loopCount:]...)
			loopCount += 1
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

	// print map
	// for _, row := range galixies {
	// 	fmt.Println(row)
	// }

	distances := map[string]int{}
	total := 0

	for num1, coords1 := range positions {
		for num2, coords2 := range positions {
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
