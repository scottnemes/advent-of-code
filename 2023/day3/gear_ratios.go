package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func isSymbol(r rune) bool {
	return unicode.IsSymbol(r) || (unicode.IsPunct(r) && string(r) != ".")
}

func solution() (int, int) {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1, -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	
	var engine [][]string
	rows := 0
	cols := 0
	symbols := map[int][]int{}
	nums := map[int]map[int][]int{}
	num := ""
	j := []int{} // used to track indexes of numbers

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		num = ""
		j = []int{}
		for i, c := range line {
			row = append(row, string(c))
			// parse numbers
			if unicode.IsNumber(c) {
				num += string(c)
				j = append(j, i) // keep track of the columns/indexes that the number takes up
			} else if num != "" {
				// number stopped, store it and restart
				n, _ := strconv.Atoi(num)
				// add one before and one after the number to catch neighbors/diagonals
				j = append(j, j[0]-1, j[len(j)-1]+1)
				if nums[n] == nil {
					nums[n] = map[int][]int{}
				}
				nums[n][rows] = append(nums[n][rows], j...)
				num = ""
				j = []int{}
			}
			// parse symbols
			if isSymbol(c) {
				symbols[rows] = append(symbols[rows], i)
			}
		}
		// check if there was a number at the end of the line
		if num != "" {
			n, _ := strconv.Atoi(num)
			// add one before and one after the number to catch neighbors/diagonals
			j = append(j, j[0]-1, j[len(j)-1]+1)
			if nums[n] == nil {
				nums[n] = map[int][]int{}
			}
			nums[n][rows] = j
		}

		engine = append(engine, row)
		cols = max(cols, len(row))
		rows += 1
	}

	p1Total := 0
	p2Total := 0

	// fmt.Println(nums)
	// fmt.Println(symbols)
	for k, v := range nums {
		// k = number
		// v = map of col/row
		for i, j := range v {
			// i = row number
			// j = []int of cols
			for _, g := range j {
				// g = cols
				if slices.Contains(symbols[i-1], g) || 
				slices.Contains(symbols[i], g) || 
				slices.Contains(symbols[i+1], g) {
					p1Total += k
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return p1Total, p2Total
}

func main() {
	// part 1
	p1Total, p2Total := solution()
	fmt.Println("Part 1:", p1Total)
	fmt.Println("Part 2:", p2Total)
}
