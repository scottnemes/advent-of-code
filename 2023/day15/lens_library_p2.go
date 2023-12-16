package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getAsciiValue(c rune) int {
	num, _ := strconv.Atoi(fmt.Sprintf("%d", c))
	return num
}

func performOp(step string, boxes map[int][][]string) {
	box := 0
	label := ""
	op := ""
	for _, char := range step {
		if string(char) == "-" || string(char) == "=" {
			op = string(char)
			break
		}
		box += getAsciiValue(char)
		box *= 17
		box = box % 256
		label += string(char)
	}
	boxSize := len(boxes[box])
	lensNum := strings.Split(step, op)[1]
	if op == "-" && boxSize > 0 {
		for i, lens := range boxes[box] {
			if lens[0] == label {
				if len(boxes[box]) == 1 {
					boxes[box] = [][]string{}
					break
				}
				new := append([][]string{}, boxes[box][:i]...)
				boxes[box] = append(new, boxes[box][i+1:]...)
				break
			}
		}
	} else if op == "=" {
		if boxSize == 0 {
			boxes[box] = append([][]string{}, []string{label, lensNum})
		} else {
			replacedLens := false
			for i, lens := range boxes[box] {
				if lens[0] == label {
					boxes[box][i][1] = lensNum
					replacedLens = true
					break
				}
			}
			if !replacedLens {
				boxes[box] = append(boxes[box], []string{label, lensNum})
			}
		}
	}
}

func solution() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)
	total := 0
	boxes := map[int][][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		steps := strings.Split(line, ",")
		for _, step := range steps {
			performOp(step, boxes)
		}
	}

	focusingPower := 0
	focalLength := 0
	for box, lenses := range boxes {
		if len(lenses) == 0 {
			continue
		}
		for i, lens := range lenses {
			focalLength, _ = strconv.Atoi(lens[1])
			focusingPower = (box + 1) * (i + 1) * focalLength
			total += focusingPower
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
