package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	order := make(map[string][]string)
	updates := [][]string{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			order[pages[1]] = append(order[pages[1]], pages[0])
		} else if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	total := 0
	for _, update := range updates {
		seen := []string{}
		valid := true
		for _, page := range update {
			reqs, ok := order[page]
			if ok {
				for _, req := range reqs {
					if slices.Contains(update, req) && !slices.Contains(seen, req) {
						valid = false
						break
					}
				}
			}
			if !valid {
				break
			}
			seen = append(seen, page)
		}
		if valid {
			middle, _ := strconv.Atoi(update[len(update)/2])
			total += middle
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
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

	order := make(map[string][]string)
	updates := [][]string{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			order[pages[1]] = append(order[pages[1]], pages[0])
		} else if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	invalid_updates := [][]string{}
	for _, update := range updates {
		seen := []string{}
		valid := true
		for _, page := range update {
			reqs, ok := order[page]
			if ok {
				for _, req := range reqs {
					if slices.Contains(update, req) && !slices.Contains(seen, req) {
						valid = false
						break
					}
				}
			}
			if !valid {
				invalid_updates = append(invalid_updates, update)
				break
			}
			seen = append(seen, page)
		}
	}

	total := 0

	for _, update := range invalid_updates {
		finished := false
		for !finished {
			seen := []string{}
			changed := false
			for current_idx, page := range update {
				reqs, ok := order[page]
				if ok {
					for _, req := range reqs {
						if slices.Contains(update, req) && !slices.Contains(seen, req) {
							req_idx := -1
							for i, c := range update {
								if c == req {
									req_idx = i
								}
							}
							update = slices.Delete(update, req_idx, req_idx+1)
							update = slices.Insert(update, current_idx, req)
							changed = true
							break
						}
						if changed {
							break
						}
					}
				}
				if changed {
					break
				}
				seen = append(seen, page)
			}
			if !changed {
				middle, _ := strconv.Atoi(update[len(update)/2])
				total += middle
				finished = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
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
