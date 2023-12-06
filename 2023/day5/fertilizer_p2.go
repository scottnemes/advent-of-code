package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// input map: # # # (destination range start, the source range start, and the range length)
// Any source numbers that aren't mapped correspond to the same destination

func addToMap(m map[int]map[string]int, key int, vals []string) {
	dstStart, _ := strconv.Atoi(vals[0])
	srcStart, _ := strconv.Atoi(vals[1])
	rng, _ := strconv.Atoi(vals[2])
	// key is the source start - source end
	// dest can be calculated by substracting the lookup value from the srcStart, and adding that to the dstStart
	entry := map[string]int{}
	entry["srcStart"] = srcStart
	entry["srcEnd"] = srcStart + rng - 1
	entry["dstStart"] = dstStart
	m[key] = entry
}

func lookupValue(m map[int]map[string]int, v int) int {
	srcStart := -1
	srcEnd := -1
	dest := -1
	for _, d := range m {
		// parse start/end source range from map key
		srcStart = d["srcStart"]
		srcEnd = d["srcEnd"]
		// if our value (v) is within this range, use it
		if srcStart <= v && v <= srcEnd {
			dest = d["dstStart"] + (v - srcStart)
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
	seedToSoilMap := map[int]map[string]int{}
	soilToFertilizer := map[int]map[string]int{}
	fertilizerToWater := map[int]map[string]int{}
	waterToLight := map[int]map[string]int{}
	lightToTemperature := map[int]map[string]int{}
	temperatureToHumidity := map[int]map[string]int{}
	humidityToLocation := map[int]map[string]int{}
	
	x := 0
	for scanner.Scan() {
		x += 1
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
			addToMap(seedToSoilMap, x, values)
		case "soil-to-fertilizer":
			addToMap(soilToFertilizer, x, values)
		case "fertilizer-to-water":
			addToMap(fertilizerToWater, x, values)
		case "water-to-light":
			addToMap(waterToLight, x, values)
		case "light-to-temperature":
			addToMap(lightToTemperature, x, values)
		case "temperature-to-humidity":
			addToMap(temperatureToHumidity, x, values)
		case "humidity-to-location":
			addToMap(humidityToLocation, x, values)
		}
	}

	closest := 999999999999999999

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for _, seedPair := range seeds {
		seedStart := seedPair[0]
		seedRange := seedPair[1]
		for i := 0; i < seedRange; i += 1 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				defer mutex.Unlock()
				seed := seedStart + i
				soil := lookupValue(seedToSoilMap, seed)
				fert := lookupValue(soilToFertilizer, soil)
				water := lookupValue(fertilizerToWater, fert)
				light := lookupValue(waterToLight, water)
				temp := lookupValue(lightToTemperature, light)
				hum := lookupValue(temperatureToHumidity, temp)
				loc := lookupValue(humidityToLocation, hum)
				mutex.Lock()
				closest = min(closest, loc)
			}(i)
		}
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	return closest
}

func main() {
	total := solution()
	fmt.Println("Part 1:", total)
}
