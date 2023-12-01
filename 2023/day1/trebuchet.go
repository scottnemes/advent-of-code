package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func solution(part int) int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	nums := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	total := 0

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()

		l := 0
		r := len(line) - 1

		type value struct {
			num string
			idx int
		}

		v1 := value{num: "", idx: -1}
		v2 := value{num: "", idx: -1}

		for l < len(line) && r >= 0 {
			// if we haven't found a left-most or right-most number, and this is a number, use it
			if v1.num == "" && unicode.IsDigit(rune(line[l])) {
				v1.num = string(line[l])
				v1.idx = l
			}
			if v2.num == "" && unicode.IsDigit(rune(line[r])) {
				v2.num = string(line[r])
				v2.idx = r
			}
			// if we have the left and right-most numbers, stop looking
			if v1.num != "" && v2.num != "" {
				break
			}
			l += 1
			r -= 1
		}
		if part == 2 {
			// check the line for written out numbers
			for k, v := range nums {
				// if the number is not in the line at all, skip it
				if strings.Index(line, k) == -1 {
					continue
				}
				// if the index of the written out number is before (for left-most) or after (for right-most), use that instead
				if strings.Index(line, k) >= 0 && (strings.Index(line, k) < v1.idx || v1.idx == -1) {
					v1.num = v
					v1.idx = strings.Index(line, k)
				}
				if strings.LastIndex(line, k) >= 0 && (strings.LastIndex(line, k) > v2.idx || v2.idx == -1) {
					v2.num = v
					v2.idx = strings.LastIndex(line, k)
				}
			}
		}
		ivalue, _ := strconv.Atoi(fmt.Sprintf("%s%s", v1.num, v2.num))
		total += ivalue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return total
}

func main() {
	// part 1
	total := solution(1)
	fmt.Println("Part 1:", total)

	// part 2
	total = solution(2)
	fmt.Println("Part 2:", total)
}
