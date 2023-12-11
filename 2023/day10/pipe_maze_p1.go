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

func solution() int {
	inputs, err := os.Open("input.txt")
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

	y := 0
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
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
		}
		y += 1
	}

	currentPipe := mazeStart
	nextPipe := ""
	nextX := 0
	nextY := 0
	moves := 0
	moveHistory := map[int]string{}
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
				fmt.Println("(No-pipe) Tried Moving", direction, "from", currentPipe, "to", nextPipe)
				continue
			}
			if maze[nextPipe]["visited"] == 1 {
				// already visited this pipe, move on
				fmt.Println("(Already Visited) Tried Moving", direction, "from", currentPipe, "to", nextPipe)
				continue
			}
			// check if the nextPipe has an opening in the direction we're heading
			validDirection := false
			for _, d := range pieces[pipeIntToString[maze[nextPipe]["type"]]] {
				if d == directionOpposite[direction] {
					validDirection = true
					//fmt.Println(pipeIntToString[maze[nextPipe]["type"]], "Found matching direction:", "Current:", currentPipe, "Next:", nextPipe, "Direction:", direction, "Opp Direction:", directionOpposite[direction], "d:", d)					break
				}
				fmt.Println("No available direction found.")
			}
			if !validDirection {
				fmt.Println("(Invalid Direction) Tried Moving", direction, "from", currentPipe, "to", nextPipe)
				continue
			}
			if pipeIntToString[maze[nextPipe]["type"]] == "S" {
				// if we've made it back to the start, stop
				fmt.Println("Done")
				done = true
				break
			}
			fmt.Println("Moving", direction, "from", currentPipe, "to", nextPipe)
			currentPipe = nextPipe
			moves += 1
			moveHistory[moves] = nextPipe
			// we've made a move, so move on to the next round
			loop = true
			break
		}
		if done {
			break
		}
	}
	// fmt.Println(moveHistory)
	// fmt.Println(currentPipe, nextPipe)
	return (moves + 1) / 2
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
