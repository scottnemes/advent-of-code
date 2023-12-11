package main

import (
	"bufio"
	"fmt"
	"os"
)

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

func checkDirection(x int, y int, d string, m [][]string) bool {
	found := false
	if d == "up" {
		for i := y - 1; i > 0; i -= 1 {
			if m[i][x] == "X" {
				found = true
				break
			}
		}
	}
	if d == "down" {
		for i := y + 1; i < len(m); i += 1 {
			if m[i][x] == "X" {
				found = true
				break
			}
		}
	}
	if d == "left" {
		found = false
		for i := x - 1; i > 0; i -= 1 {
			if m[y][i] == "X" {
				found = true
				break
			}
		}
	}
	if d == "right" {
		found = false
		for i := x + 1; i < len(m[y]); i += 1 {
			if m[y][i] == "X" {
				found = true
				break
			}
		}
	}
	if found {
		foundBound := false
		for x2 := x; x2 < len(m[y]); x2 += 1 {
			for y2 := y; y2 >= 0; y2 -= 1 {
				if m[y2][x2] == "X" {
					foundBound = true
					break
				}
			}
			if !foundBound {
				return false
			}
			for y2 := y; y2 < len(m); y2 += 1 {
				if m[y2][x2] == "X" {
					foundBound = true
					break
				}
			}
			if !foundBound {
				return false
			}
		}
		foundBound = false
		for x2 := x; x2 > 0; x2 -= 1 {
			for y2 := y; y2 >= 0; y2 -= 1 {
				if m[y2][x2] == "X" {
					foundBound = true
					break
				}
			}
			if !foundBound {
				return false
			}
			for y2 := y; y2 < len(m); y2 += 1 {
				if m[y2][x2] == "X" {
					foundBound = true
					break
				}
			}
			if !foundBound {
				return false
			}
		}
	}

	return found
}

func solution() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	maze := map[string]map[string]int{}
	mazeStart := ""

	pipeIntToString := map[int]string{
		0: "|",
		1: "-",
		2: "L",
		3: "J",
		4: "7",
		5: "F",
		6: "S",
	}

	pipeStringToInt := map[string]int{
		"|": 0,
		"-": 1,
		"L": 2,
		"J": 3,
		"7": 4,
		"F": 5,
		"S": 6,
		".": 7,
	}

	pieces := map[string][]string{
		"|": {"N", "S"},
		"-": {"E", "W"},
		"L": {"N", "E"},
		"J": {"N", "W"},
		"7": {"S", "W"},
		"F": {"S", "E"},
		"S": {"N", "S", "E", "W"},
	}

	directions := map[string]int{
		"N": -1,
		"S": 1,
		"W": -1,
		"E": 1,
	}

	directionOpposite := map[string]string{
		"N": "S",
		"S": "N",
		"E": "W",
		"W": "E",
	}

	moveHistory := [][]string{}

	y := 0
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		emptyLine := []string{}
		for x, p := range line {
			name := fmt.Sprintf("%d-%d", x, y)
			pipe := map[string]int{
				"x":       x,
				"y":       y,
				"type":    pipeStringToInt[string(p)],
				"visited": 0,
			}
			if string(p) == "S" {
				mazeStart = name
			}
			maze[name] = pipe
			emptyLine = append(emptyLine, ".")
		}
		moveHistory = append(moveHistory, emptyLine)
		y += 1
	}

	currentPipe := mazeStart
	nextPipe := ""
	nextX := 0
	nextY := 0
	moves := 0
	done := false
	loop := true
	for loop {
		maze[currentPipe]["visited"] = 1
		loop = false
		for _, direction := range pieces[pipeIntToString[maze[currentPipe]["type"]]] {
			nextX = maze[currentPipe]["x"]
			nextY = maze[currentPipe]["y"]
			if direction == "N" || direction == "S" {
				// +1 if S, -1 if N
				nextY += directions[direction]
			} else {
				// +1 if E, -1 if W
				nextX += directions[direction]
			}
			nextPipe = fmt.Sprintf("%d-%d", nextX, nextY)
			if pipeIntToString[maze[nextPipe]["type"]] == "." {
				// no pipe in this direction, move on
				//fmt.Println("(No-pipe) Tried Moving", direction, "from", currentPipe, "to", nextPipe)
				continue
			}
			if maze[nextPipe]["visited"] == 1 {
				// already visited this pipe, move on
				//fmt.Println("(Already Visited) Tried Moving", direction, "from", currentPipe, "to", nextPipe)
				continue
			}
			// check if the nextPipe has an opening in the direction we're heading
			validDirection := false
			for _, d := range pieces[pipeIntToString[maze[nextPipe]["type"]]] {
				if d == directionOpposite[direction] {
					validDirection = true
					//fmt.Println(pipeIntToString[maze[nextPipe]["type"]], "Found matching direction:", "Current:", currentPipe, "Next:", nextPipe, "Direction:", direction, "Opp Direction:", directionOpposite[direction], "d:", d)					break
				}
				//fmt.Println("No available direction found.")
			}
			if !validDirection {
				//fmt.Println("(Invalid Direction) Tried Moving", direction, "from", currentPipe, "to", nextPipe)
				continue
			}
			if pipeIntToString[maze[nextPipe]["type"]] == "S" {
				// if we've made it back to the start, stop
				fmt.Println("Done")
				done = true
				break
			}
			//fmt.Println("Moving", direction, "from", currentPipe, "to", nextPipe)
			currentPipe = nextPipe
			moves += 1
			moveHistory[maze[nextPipe]["y"]][maze[nextPipe]["x"]] = "X"
			// we've made a move, so move on to the next round
			loop = true
			break
		}
		if done {
			break
		}
	}

	for _, row := range moveHistory {
		fmt.Println(row)
	}

	total := 0
	found := false
	for y, row := range moveHistory {
		if y < 1 || y > len(moveHistory)-2 {
			continue
		}
		for x, v := range row {
			if x < 1 || x > len(row)-2 || v == "X" {
				continue
			}
			// check up
			found = checkDirection(x, y, "up", moveHistory)
			if !found {
				continue
			}
			// check down
			found = checkDirection(x, y, "down", moveHistory)
			if !found {
				continue
			}
			// check left
			found = checkDirection(x, y, "left", moveHistory)
			if !found {
				continue
			}
			// check right
			found = checkDirection(x, y, "right", moveHistory)
			if !found {
				continue
			}
			total += 1
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
