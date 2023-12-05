package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// input map: # # # (destination range start, the source range start, and the range length)
// Any source numbers that aren't mapped correspond to the same destination

func addToMap(m map[string]int, vals []string) {
	dstStart, _ := strconv.Atoi(vals[0])
	srcStart, _ := strconv.Atoi(vals[1])
	rng, _ := strconv.Atoi(vals[2])
	// key is the source start - source end
	// dest can be calculated by substracting the lookup value from the srcStart, and adding that to the dstStart
	m[fmt.Sprintf("%d-%d", srcStart, srcStart + rng - 1)] = dstStart
}

func lookupValue(m map[string]int, v int) int {
	srcStart := -1
	srcEnd := -1
	dest := -1
	for rng, d := range m {
		// parse start/end source range from map key
		srcStart, _ = strconv.Atoi(strings.Split(rng, "-")[0])
		srcEnd, _ = strconv.Atoi(strings.Split(rng, "-")[1])
		// if our value (v) is within this range, use it
		if srcStart <= v && v <= srcEnd {
			dest = d + (v - srcStart)
		}
	}
	// if we found no mapping for this value, use the value itself
	if dest == -1 {
		dest = v
	}
	return dest
}

func solution() (int) {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	prefix := ""
	values := []string{}
	seeds := [][]int{}
	seedToSoilMap := map[string]int{}
	soilToFertilizer := map[string]int{}
	fertilizerToWater := map[string]int{}
	waterToLight := map[string]int{}
	lightToTemperature := map[string]int{}
	temperatureToHumidity := map[string]int{}
	humidityToLocation := map[string]int{}
	
	loc := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// if this line is not a mapping range, determine which map it is
		if _, err := strconv.Atoi(string(line[0])); err != nil {
			prefix = strings.Split(line, " ")[0]
			if prefix == "seeds:" {
				x := 0
				seedPair := []int{}
				for _, i := range strings.Split(line, " ")[1:] {
					num, _ := strconv.Atoi(i)
					seedPair = append(seedPair, num)
					// first value is the seed start, second is range, restart after that
					if x == 1 {
						seeds = append(seeds, seedPair)
						x = 0
						seedPair = []int{}
					} else {
						x += 1
					}
				}
			}
			continue
		}
		values = strings.Split(line, " ")
		switch prefix {
		case "seed-to-soil":
			addToMap(seedToSoilMap, values)
		case "soil-to-fertilizer":
			addToMap(soilToFertilizer, values)
		case "fertilizer-to-water":
			addToMap(fertilizerToWater, values)
		case "water-to-light":
			addToMap(waterToLight, values)
		case "light-to-temperature":
			addToMap(lightToTemperature, values)
		case "temperature-to-humidity":
			addToMap(temperatureToHumidity, values)
		case "humidity-to-location":
			addToMap(humidityToLocation, values)
		}
	}

	soil := -1
	fert := -1
	water := -1
	light := -1
	temp := -1
	hum := -1

	closest := []int{}

	seed := -1
	seedStart := -1
	seedRange := -1

	for _, seedPair := range seeds {
		seedStart = seedPair[0]
		seedRange = seedPair[1]
		for i := 0; i < seedRange; i += 1 {
			seed = seedStart + i
			soil = lookupValue(seedToSoilMap, seed)
			fert = lookupValue(soilToFertilizer, soil)
			water = lookupValue(fertilizerToWater, fert)
			light = lookupValue(waterToLight, water)
			temp = lookupValue(lightToTemperature, light)
			hum = lookupValue(temperatureToHumidity, temp)
			loc = lookupValue(humidityToLocation, hum)
			closest = append(closest, loc)
			// fmt.Println("Seed:", seed)
			// fmt.Println("Soil:", soil)
			// fmt.Println("Fert:", fert)
			// fmt.Println("Water:", water)
			// fmt.Println("Light:", light)
			// fmt.Println("Temp:", temp)
			// fmt.Println("Hum:", hum)
			// fmt.Println("Loc:", loc)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return slices.Min(closest)
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
