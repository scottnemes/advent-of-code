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

type symbol struct {
	val string
	row int
	col int
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
	symbols := map[int][]symbol{}
	nums := map[int]map[int][]int{}
	num := ""
	numIdxs := []int{} // used to track indexes of numbers

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		num = ""
		numIdxs = []int{}
		for i, c := range line {
			row = append(row, string(c))
			// parse numbers
			if unicode.IsNumber(c) {
				num += string(c)
				numIdxs = append(numIdxs, i) // keep track of the columns/indexes that the number takes up
			} else if num != "" {
				// number stopped, store it and restart
				n, _ := strconv.Atoi(num)
				// add one before and one after the number to catch neighbors/diagonals
				numIdxs = append(numIdxs, numIdxs[0]-1, numIdxs[len(numIdxs)-1]+1)
				if nums[n] == nil {
					nums[n] = map[int][]int{}
				}
				nums[n][rows] = append(nums[n][rows], numIdxs...)
				num = ""
				numIdxs = []int{}
			}
			// parse symbols
			if isSymbol(c) {
				symbols[rows] = append(symbols[rows], symbol{val: string(c), row: rows, col: i})
			}
		}
		// check if there was a number at the end of the line
		if num != "" {
			n, _ := strconv.Atoi(num)
			// add one before and one after the number to catch neighbors/diagonals
			numIdxs = append(numIdxs, numIdxs[0]-1, numIdxs[len(numIdxs)-1]+1)
			if nums[n] == nil {
				nums[n] = map[int][]int{}
			}
			nums[n][rows] = numIdxs
		}

		engine = append(engine, row)
		rows += 1
	}

	p1Total := 0
	p2Total := 0

	// part 1
	cols1 := []int{}
	cols2 := []int{}
	cols3 := []int{}
	for k, v := range nums {
		// k = number
		// v = map of col/row
		for i, j := range v {
			cols1 = []int{}
			cols2 = []int{}
			cols3 = []int{}
			// i = row number
			// j = []int of cols
			for _, g := range j {
				// g = cols
				for _, s := range symbols[i-1] {
					cols1 = append(cols1, s.col)
				}
				for _, s := range symbols[i] {
					cols2 = append(cols2, s.col)
				}
				for _, s := range symbols[i+1] {
					cols3 = append(cols3, s.col)
				}
				if slices.Contains(cols1, g) || 
				slices.Contains(cols2, g) || 
				slices.Contains(cols3, g) {
					p1Total += k
				}
			}
		}
	}

	// part 2
	adjNums := []int{}
	for _, syms := range symbols {
		for _, sym := range syms {
			// skip if this is not a *
			if sym.val != "*" {
				continue
			}
			for num, v := range nums {
				// v = map of col/row
				for numRow, j := range v {
					// j = []int of cols
					// if the number is not in an adjacent row to the symbol, skip it
					if numRow < (sym.row - 1) || numRow > (sym.row + 1) {
						continue
					}
					for _, numCol := range j {
						if sym.col == numCol {
							adjNums = append(adjNums, num)
							break
						}
					}
				}
			}
			// if the * has exactly two neighboring numbers, use it
			if len(adjNums) == 2 {
				p2Total += (adjNums[0] * adjNums[1])
			}
			adjNums = []int{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return p1Total, p2Total
}

func main() {
	p1Total, p2Total := solution()
	fmt.Println("Part 1:", p1Total)
	fmt.Println("Part 2:", p2Total)
}
