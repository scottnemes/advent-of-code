package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	parts := map[string]bool{}
	workflows := map[string][]string{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// parse workflows
		if string(line[0]) != "{" {
			temp := strings.Split(line, "{")
			name := temp[0]
			steps := strings.Split(temp[1][:len(temp[1])-1], ",")
			workflows[name] = steps
		} else if string(line[0]) == "{" {
			// parse parts
			parts[line[1:len(line)-1]] = false
		}
	}
	for part, _ := range parts {
		temp := strings.Split(part, ",")
		values := map[string]int{}
		for _, val := range temp {
			k := strings.Split(val, "=")[0]
			v, _ := strconv.Atoi(strings.Split(val, "=")[1])
			switch k {
			case "x":
				values["x"] = v
			case "m":
				values["m"] = v
			case "a":
				values["a"] = v
			case "s":
				values["s"] = v
			}
		}
		workflow := workflows["in"]
		partFinished := false
		for true {
			partFinished = false
			for _, step := range workflow {
				check := ""
				next := ""
				// step does not contain a comparison check
				if len(strings.Split(step, ":")) == 1 {
					next = strings.Split(step, ":")[0]
					if next == "A" {
						parts[part] = true
						partFinished = true
					} else if next == "R" {
						parts[part] = false
						partFinished = true
					} else {
						workflow = workflows[next]
					}
					break
				}
				check = strings.Split(step, ":")[0]
				next = strings.Split(step, ":")[1]
				// {x=787,m=2655,a=1222,s=2876}
				// px{a<2006:qkq,m>2090:A,rfg}
				// mnt{m<1258:gzs,cr}
				operator := ""
				if strings.Contains(check, "<") {
					operator = "<"
				} else {
					operator = ">"
				}
				catCount := values[strings.Split(check, operator)[0]]
				val, _ := strconv.Atoi(strings.Split(check, operator)[1])
				pass := false
				if operator == "<" {
					if catCount < val {
						pass = true
					}
				} else {
					if catCount > val {
						pass = true
					}
				}
				// passed check, go to target workflow for given check
				if pass {
					if next == "A" {
						parts[part] = true
						partFinished = true
					} else if next == "R" {
						parts[part] = false
						partFinished = true
					} else {
						workflow = workflows[next]
					}
					break
				}
				// check did not pass, go on to the next step for this same workflow
			}
			// if part is accepted or rejected already, stop checking
			if partFinished {
				break
			}
		}
	}

	total := 0
	sub := 0
	for vals, accepted := range parts {
		if !accepted {
			continue
		}
		temp := strings.Split(vals, ",")
		sub = 0
		for _, val := range temp {
			v, _ := strconv.Atoi(strings.Split(val, "=")[1])
			sub += v
		}
		total += sub
	}

	return total
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
