package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	guard_map := [][]string{}
	pos := make([]int, 2)
	facing := "N"

	scanner := bufio.NewScanner(inputs)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		r := []string{}
		for x, char := range line {
			c := string(char)
			r = append(r, c)
			if c == "^" {
				pos[0], pos[1] = x, y
			}
		}
		guard_map = append(guard_map, r)
		y += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	max_x := len(guard_map[0]) - 1
	max_y := len(guard_map) - 1

	total := 1 // includes start

	done := false
	for !done {
		x := pos[0]
		y := pos[1]
		switch facing {
		case "N":
			if y-1 < 0 {
				// edge of map
				done = true
				break
			}
			if guard_map[y-1][x] == "#" {
				// blocked
				facing = "E"
				break
			}
			pos[1] -= 1
			if guard_map[y][x] != "X" {
				guard_map[y][x] = "X"
				total += 1
			}
		case "S":
			if y+1 > max_y {
				// edge of map
				done = true
				break
			}
			if guard_map[y+1][x] == "#" {
				// blocked
				facing = "W"
				break
			}
			pos[1] += 1
			if guard_map[y][x] != "X" {
				guard_map[y][x] = "X"
				total += 1
			}
		case "E":
			if x+1 > max_x {
				// edge of map
				done = true
				break
			}
			if guard_map[y][x+1] == "#" {
				// blocked
				facing = "S"
				break
			}
			pos[0] += 1
			if guard_map[y][x] != "X" {
				guard_map[y][x] = "X"
				total += 1
			}
		case "W":
			if x-1 < 0 {
				// edge of map
				done = true
				break
			}
			if guard_map[y][x-1] == "#" {
				// blocked
				facing = "N"
				break
			}
			pos[0] -= 1
			if guard_map[y][x] != "X" {
				guard_map[y][x] = "X"
				total += 1
			}
		}
	}

	return total
}

func solution2() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	guard_map := [][]string{}
	pos := make([]int, 2)
	facing := "N"

	scanner := bufio.NewScanner(inputs)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		r := []string{}
		for x, char := range line {
			c := string(char)
			r = append(r, c)
			if c == "^" {
				pos[0], pos[1] = x, y
			}
		}
		guard_map = append(guard_map, r)
		y += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	start_pos := make([]int, 2)
	copy(start_pos, pos)

	max_x := len(guard_map[0]) - 1
	max_y := len(guard_map) - 1

	visited_coords := [][]int{}

	done := false
	for !done {
		x := pos[0]
		y := pos[1]
		switch facing {
		case "N":
			if y-1 < 0 {
				// edge of map
				done = true
				break
			}
			if guard_map[y-1][x] == "#" {
				// blocked
				facing = "E"
				break
			}
			pos[1] -= 1
			if guard_map[y-1][x] != "X" {
				guard_map[y-1][x] = "X"
				c := make([]int, 2)
				copy(c, pos)
				visited_coords = append(visited_coords, c)
			}
		case "S":
			if y+1 > max_y {
				// edge of map
				done = true
				break
			}
			if guard_map[y+1][x] == "#" {
				// blocked
				facing = "W"
				break
			}
			pos[1] += 1
			if guard_map[y+1][x] != "X" {
				guard_map[y+1][x] = "X"
				c := make([]int, 2)
				copy(c, pos)
				visited_coords = append(visited_coords, c)
			}
		case "E":
			if x+1 > max_x {
				// edge of map
				done = true
				break
			}
			if guard_map[y][x+1] == "#" {
				// blocked
				facing = "S"
				break
			}
			pos[0] += 1
			if guard_map[y][x+1] != "X" {
				guard_map[y][x+1] = "X"
				c := make([]int, 2)
				copy(c, pos)
				visited_coords = append(visited_coords, c)
			}
		case "W":
			if x-1 < 0 {
				// edge of map
				done = true
				break
			}
			if guard_map[y][x-1] == "#" {
				// blocked
				facing = "N"
				break
			}
			pos[0] -= 1
			if guard_map[y][x-1] != "X" {
				guard_map[y][x-1] = "X"
				c := make([]int, 2)
				copy(c, pos)
				visited_coords = append(visited_coords, c)
			}
		}
	}

	guard_map_copy := make([][]string, len(guard_map))
	for i := range guard_map {
		guard_map_copy[i] = make([]string, len(guard_map[i]))
		copy(guard_map_copy[i], guard_map[i])
	}

	total := 0

	// 3,6
	// 6,7
	// 7,7
	// 1,8
	// 3,8
	// 7,9

	// visited_coords = [][]int{{7, 8}}

	for _, coords := range visited_coords {
		// added_block := fmt.Sprintf("%d-%d", coords[0], coords[1])
		copy(pos, start_pos)
		guard_map := make([][]string, len(guard_map_copy))
		for i := range guard_map_copy {
			guard_map[i] = make([]string, len(guard_map_copy[i]))
			copy(guard_map[i], guard_map_copy[i])
		}
		guard_map[coords[1]][coords[0]] = "#"
		done = false
		facing = "N"
		// visited_blocks := []string{fmt.Sprintf("%d-%d-%s", pos[0], pos[1], facing)}
		visited_blocks := []string{}
		fmt.Println(coords)
		for !done {
			x := pos[0]
			y := pos[1]
			switch facing {
			case "N":
				next_x := x
				next_y := y - 1
				if next_y < 0 {
					// edge of map
					done = true
					break
				}
				tile := fmt.Sprintf("%d-%d-%s", next_x, next_y, facing)
				if slices.Contains(visited_blocks, tile) {
					done = true
					total += 1
					break
				}
				if guard_map[next_y][next_x] == "#" {
					// blocked
					facing = "E"
					break
				}
				visited_blocks = append(visited_blocks, tile)
				pos[1] = next_y
			case "S":
				next_x := x
				next_y := y + 1
				if next_y > max_y {
					// edge of map
					done = true
					break
				}
				tile := fmt.Sprintf("%d-%d-%s", next_x, next_y, facing)
				if slices.Contains(visited_blocks, tile) {
					done = true
					total += 1
					break
				}
				if guard_map[next_y][next_x] == "#" {
					// blocked
					facing = "W"
					break
				}
				visited_blocks = append(visited_blocks, tile)
				pos[1] = next_y
			case "E":
				next_x := x + 1
				next_y := y
				if next_x > max_x {
					// edge of map
					done = true
					break
				}
				tile := fmt.Sprintf("%d-%d-%s", next_x, next_y, facing)
				if slices.Contains(visited_blocks, tile) {
					done = true
					total += 1
					break
				}
				if guard_map[next_y][next_x] == "#" {
					// blocked
					facing = "S"
					break
				}
				visited_blocks = append(visited_blocks, tile)
				pos[0] = next_x
			case "W":
				next_x := x - 1
				next_y := y
				if next_x < 0 {
					// edge of map
					done = true
					break
				}
				tile := fmt.Sprintf("%d-%d-%s", next_x, next_y, facing)
				if slices.Contains(visited_blocks, tile) {
					done = true
					total += 1
					break
				}
				if guard_map[next_y][next_x] == "#" {
					// blocked
					facing = "N"
					break
				}
				visited_blocks = append(visited_blocks, tile)
				pos[0] = next_x
			}
		}
	}

	return total
}

func main() {
	// part 1
	total := solution1()
	fmt.Println("Part 1:", total)

	// part 2
	total = solution2()
	fmt.Println("Part 2:", total)
}
