package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PulseModule struct {
	mType        string
	state        bool
	history      map[string]string
	destinations []string
}

func sendPulse(moduleType string, pulseModule string, m map[string]PulseModule) {
	switch moduleType {
	case "broadcaster":
		processBroadcaster(pulseModule, m)
	case "flip-flop":
		processFlipFlop(pulseModule, m)
	case "conjunction":
		processConjunction(pulseModule, m)
	}
}

func processBroadcaster(pulseModule string, m map[string]PulseModule) {

}

func processFlipFlop(pulseModule string, m map[string]PulseModule) {

}

func processConjunction(pulseModule string, m map[string]PulseModule) {

}

func solution() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	pulseModules := map[string]PulseModule{}
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, " -> ")
		pm := PulseModule{}
		mType := ""
		name := ""
		if string(temp[0][0]) == "%" || string(temp[0][0]) == "&" {
			name = temp[0][1:]
			mType = string(temp[0][0])
		} else {
			name = temp[0]
			mType = "broadcaster"
		}
		switch mType {
		case "broadcaster":
			pm.mType = "broadcaster"
		case "%":
			pm.mType = "flip-flop"
			pm.state = false
		case "&":
			pm.mType = "conjunction"
		}
		pm.destinations = strings.Split(temp[1], ", ")
		pulseModules[name] = pm
	}

	fmt.Println(pulseModules)

	return -1
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
