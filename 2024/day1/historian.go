package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type value struct {
	num string
	idx int
}

func solution1() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	locations_one := []float64{}
	locations_two := []float64{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		locations := strings.Split(line, "   ")
		one, _ := strconv.Atoi(locations[0])
		two, _ := strconv.Atoi(locations[1])
		locations_one = append(locations_one, float64(one))
		locations_two = append(locations_two, float64(two))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	sort.Slice(locations_one, func(i, j int) bool {
		return locations_one[i] > locations_one[j]
	})

	sort.Slice(locations_two, func(i, j int) bool {
		return locations_two[i] > locations_two[j]
	})

	total := 0

	for i := range locations_one {
		total += int(math.Abs(locations_one[i] - locations_two[i]))
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

	locations_one := []int{}
	locations_two := []int{}

	locations_two_freq := map[int]int{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		locations := strings.Split(line, "   ")
		one, _ := strconv.Atoi(locations[0])
		two, _ := strconv.Atoi(locations[1])
		locations_one = append(locations_one, one)
		locations_two = append(locations_two, two)
		_, ok := locations_two_freq[two]
		if ok {
			locations_two_freq[two] += 1
		} else {
			locations_two_freq[two] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	total := 0

	for _, v := range locations_one {
		_, ok := locations_two_freq[v]
		if ok {
			total += (v * locations_two_freq[v])
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
