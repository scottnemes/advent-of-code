package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time int
	distance int
}

func solution() (int) {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	race := Race{}

	for scanner.Scan() {
		line := scanner.Text()
		
		if strings.HasPrefix(line, "Time:") {
			race.time, _ = strconv.Atoi(strings.Join(strings.Fields(line)[1:], ""))
		} else {
			race.distance, _ = strconv.Atoi(strings.Join(strings.Fields(line)[1:], ""))
		}
	}

	// distance = (time - hold) * hold
	total := 0
	for hold := 0; hold < race.time; hold += 1 {
		if (race.time - hold) * hold > race.distance {
			total += 1
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
