package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	lagoon := [][]string{}
	scanner := bufio.NewScanner(inputs)
	// x := 0
	// y := 0
	rows := 0
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		count, _ := strconv.Atoi(split[1])
		// fmt.Println(lagoon)
		if dir == "U" || dir == "D" {
			rows += count
		}
		if dir == "L" || dir == "R" {
			cols += count
		}
	}
	y := rows
	x := cols
	rows *= 2
	cols *= 2
	for i := 0; i < rows; i += 1 {
		row := make([]string, cols)
		for j := 0; j < cols; j += 1 {
			row[j] = "."
		}
		lagoon = append(lagoon, row)
	}

	inputs2, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs2.Close()

	lagoon[y][x] = "#"
	scanner2 := bufio.NewScanner(inputs2)
	for scanner2.Scan() {
		line := scanner2.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		count, _ := strconv.Atoi(split[1])
		switch dir {
		case "U":
			for i := 0; i < count; i += 1 {
				lagoon[y-i][x] = "#"
			}
			y -= count
		case "D":
			for i := 0; i < count; i += 1 {
				lagoon[y+i][x] = "#"
			}
			y += count
		case "L":
			for i := 0; i < count; i += 1 {
				lagoon[y][x-i] = "#"
			}
			x -= count
		case "R":
			for i := 0; i < count; i += 1 {
				lagoon[y][x+i] = "#"
			}
			x += count
		}
	}

	compactedLagoon := [][]string{}

	for i := 0; i < len(lagoon); i += 1 {
		notEmpty := false
		for _, c := range lagoon[i] {
			if string(c) != "." {
				notEmpty = true
			}
		}
		if notEmpty {
			compactedLagoon = append(compactedLagoon, lagoon[i])
		}
	}

	startX := 9999
	endX := 0
	for _, row := range compactedLagoon {
		for j, char := range row {
			if string(char) == "#" {
				startX = min(startX, j)
				endX = max(endX, j)
			}
		}
	}

	for i := 0; i < len(compactedLagoon); i += 1 {
		compactedLagoon[i] = compactedLagoon[i][startX : endX+1]
	}

	// startX := -1
	// endX := -1

	// for i, row := range compactedLagoon {
	// 	startX = -1
	// 	endX = -1
	// 	for j, char := range row {
	// 		if string(char) == "#" {
	// 			if startX == -1 {
	// 				startX = j
	// 			} else {
	// 				endX = j
	// 			}
	// 		}
	// 	}
	// 	compactedLagoon[i] = compactedLagoon[i][startX : endX+1]
	// }

	total := 0
	//subTotal := 0
	//temp := 0
	startIdx := -1
	// endIdx := -1
	//prevChar := ""
	for i, _ := range compactedLagoon {
		startIdx = -1
		// endIdx = -1
		//subTotal = 0
		//temp = 0
		for j, char := range compactedLagoon[i] {
			// find where we have pairs of #s
			if string(char) == "#" && startIdx == -1 {
				startIdx = j
			} else if string(char) == "#" && startIdx != -1 {
				// endIdx = j
			}
			// } else if string(char) == "." && prevChar == "#" && startIdx == -1 {
			// 	startIdx = j - 1
			// }
			// fill in pairs
			// if startIdx != -1 && endIdx != -1 {
			// 	for k := startIdx; k <= endIdx; k += 1 {
			// 		compactedLagoon[i][k] = "#"
			// 	}
			// 	startIdx = -1
			// 	endIdx = -1
			// }
			// add #s to the sub total count
			// if string(char) == "#" {
			// 	subTotal += 1
			// }
			//prevChar = char
		}
		//total += subTotal
		//fmt.Println(compactedLagoon[i], subTotal)
	}

	for _, row := range compactedLagoon {
		fmt.Println(row)
	}

	// for _, row := range compactedLagoon {
	// 	total += len(row)
	// }

	for y := 0; y < len(compactedLagoon); y += 1 {
		for x := 0; x < len(compactedLagoon[y]); x += 1 {
			if string(compactedLagoon[y][x]) == "#" {
				total += 1
			}
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
