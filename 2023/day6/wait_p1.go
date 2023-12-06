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
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	races := map[int]Race{}
	times := []string{}
	distances := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			times = strings.Fields(line)[1:]
		} else {
			distances = strings.Fields(line)[1:]
		}
	}

	// build out races map
	time := -1
	distance := -1
	for i := 0; i < len(times); i += 1 {
		time, _ = strconv.Atoi(times[i])
		distance, _ = strconv.Atoi(distances[i])
		races[i] = Race{time: time, distance: distance }
	}
	// distance = (time - hold) * hold
	wins := 0
	total := -1
	for _, race := range races {
		wins = 0
		for hold := 0; hold < race.time; hold += 1 {
			if (race.time - hold) * hold > race.distance {
				wins += 1
			}
		}
		if total == -1 {
			total = wins
		} else {
			total *= wins
		}
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
