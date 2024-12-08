package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_all_options(running_total int, index int, target int, nums []int) bool {
	if running_total > target {
		return false
	}
	if index == len(nums) {
		if running_total == target {
			return true
		}
		return false
	}
	return (check_all_options(running_total+nums[index], index+1, target, nums) ||
		check_all_options(running_total*nums[index], index+1, target, nums))
}

func check_all_options2(running_total int, index int, target int, nums []int) bool {
	if running_total > target {
		return false
	}
	if index == len(nums) {
		if running_total == target {
			return true
		}
		return false
	}
	combined_number, _ := strconv.Atoi(fmt.Sprintf("%d%d", running_total, nums[index]))
	return (check_all_options2(running_total+nums[index], index+1, target, nums) ||
		check_all_options2(running_total*nums[index], index+1, target, nums) ||
		check_all_options2(combined_number, index+1, target, nums))
}

func is_valid_equation(target int, nums []int) bool {
	return check_all_options(nums[0], 1, target, nums)
}

func is_valid_equation2(target int, nums []int) bool {
	return check_all_options2(nums[0], 1, target, nums)
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
		f := strings.Fields(line)
		target, _ := strconv.Atoi(strings.Split(line, ":")[0])
		nums := make([]int, len(f)-1)
		for _, n := range f[1:] {
			n, _ := strconv.Atoi(n)
			nums = append(nums, n)
		}
		if is_valid_equation(target, nums) {
			total += target
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
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Fields(line)
		target, _ := strconv.Atoi(strings.Split(line, ":")[0])
		nums := make([]int, len(f)-1)
		for _, n := range f[1:] {
			n, _ := strconv.Atoi(n)
			nums = append(nums, n)
		}
		if is_valid_equation2(target, nums) {
			total += target
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
