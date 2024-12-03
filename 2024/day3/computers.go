package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	total := 0

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`mul\((\d+)\,(\d+)\)`)
		ops := r.FindAllString(line, -1)
		for _, op := range ops {
			r2, _ := regexp.Compile(`(\d+)`)
			nums := r2.FindAllString(op, -1)
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			total += (num1 * num2)
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

	total := 0

	scanner := bufio.NewScanner(inputs)
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`mul\((\d+)\,(\d+)\)|do\(\)|don\'t\(\)`)
		ops := r.FindAllString(line, -1)
		for _, op := range ops {
			command := strings.Split(op, "(")
			if command[0] == "do" {
				enabled = true
				continue
			} else if command[0] == "don't" {
				enabled = false
				continue
			}
			if enabled {
				r2, _ := regexp.Compile(`(\d+)`)
				nums := r2.FindAllString(op, -1)
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				total += (num1 * num2)
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
