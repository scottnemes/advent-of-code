package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type value struct {
	num string
	idx int
}

func solution1() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	puzzle_rows := []string{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		puzzle_rows = append(puzzle_rows, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	puzzle_cols := []string{}

	for x := 0; x < len(puzzle_rows[0]); x++ {
		s := ""
		for y := 0; y < len(puzzle_rows); y++ {
			s += string(puzzle_rows[y][x])
		}
		puzzle_cols = append(puzzle_cols, s)
	}

	puzzle_rows_reversed := make([]string, len(puzzle_rows))
	_ = copy(puzzle_rows_reversed, puzzle_rows)

	for i, s := range puzzle_rows_reversed {
		temp := ""
		for i := len(s) - 1; i >= 0; i-- {
			temp += string(s[i])
		}
		puzzle_rows_reversed[i] = temp
	}

	puzzle_diags := []string{}

	for y := len(puzzle_rows) - 1; y >= 0; y-- {
		s1 := ""
		s2 := ""
		for y2, x := y, 0; y2 < len(puzzle_rows) && x <= x; y2, x = y2+1, x+1 {
			s1 += string(puzzle_rows[y2][x])
			s2 += string(puzzle_rows_reversed[y2][x])
		}
		// puzzle_diags = append(puzzle_diags, s1)
		// puzzle_diags = append(puzzle_diags, s2)
	}

	for y := 0; y < len(puzzle_rows); y++ {
		s1 := ""
		s2 := ""
		for y2, x := y, y+1; y2 < len(puzzle_rows) && x < len(puzzle_rows[0]); y2, x = y2+1, x+1 {
			s1 += string(puzzle_rows[y2][x])
			s2 += string(puzzle_rows_reversed[y2][x])
		}
		puzzle_diags = append(puzzle_diags, s1)
		// puzzle_diags = append(puzzle_diags, s2)
	}

	r, _ := regexp.Compile(`XMAS|SAMX`)
	r2, _ := regexp.Compile(`XMASAMX|SAMXMAS`)

	total := 0

	for _, s := range puzzle_rows {
		matches := r.FindAllString(s, -1)
		total += len(matches)
		matches2 := r2.FindAllString(s, -1)
		total += len(matches2)
	}

	for _, s := range puzzle_cols {
		matches := r.FindAllString(s, -1)
		total += len(matches)
		matches2 := r2.FindAllString(s, -1)
		total += len(matches2)
	}

	for _, s := range puzzle_diags {
		matches := r.FindAllString(s, -1)
		total += len(matches)
		matches2 := r2.FindAllString(s, -1)
		total += len(matches2)
	}

	for _, row := range puzzle_diags {
		fmt.Println(row)
	}

	return total
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
