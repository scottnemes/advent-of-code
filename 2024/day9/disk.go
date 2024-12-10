package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func solution1() int {
	inputs, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	disk := []int{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		file_id := 0
		file_size := -1
		free_space := -1
		for j, val := range line {
			if file_size == -1 {
				file_size, _ = strconv.Atoi(string(val))
			} else {
				free_space, _ = strconv.Atoi(string(val))
			}
			if free_space != -1 || j == len(line)-1 {
				for i := 0; i < file_size; i++ {
					disk = append(disk, file_id)
				}
				for i := 0; i < free_space; i++ {
					disk = append(disk, -1)
				}
				// fmt.Println(file_id, file_size, free_space)
				file_size = -1
				free_space = -1
				file_id += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	next_empty := slices.Index(disk, -1)
	current_idx := len(disk) - 1
	for next_empty != -1 {
		for i := current_idx; i > 0; i -= 1 {
			if disk[i] == -1 {
				continue
			}
			current_idx = i
			break
		}
		val_to_move := disk[current_idx]
		disk[current_idx] = -1
		disk[next_empty] = val_to_move
		next_empty = slices.Index(disk[:current_idx], -1)
	}

	total := 0
	for i := 0; i < len(disk); i += 1 {
		if disk[i] != -1 {
			total += (i * disk[i])
		}
	}

	return total
}

func solution2() int {
	inputs, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Failed to open puzzle inputs, RIP:", err)
		return -1
	}
	defer inputs.Close()

	disk := []int{}
	free_chunks := [][]int{}
	file_id := 0
	file_sizes := map[int]int{}

	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		line := scanner.Text()
		file_size := -1
		free_space := -1
		for j, val := range line {
			if file_size == -1 {
				file_size, _ = strconv.Atoi(string(val))
			} else {
				free_space, _ = strconv.Atoi(string(val))
			}
			file_sizes[file_id] = file_size
			if free_space != -1 || j == len(line)-1 {
				for i := 0; i < file_size; i++ {
					disk = append(disk, file_id)
				}
				for i := 0; i < free_space; i++ {
					disk = append(disk, -1)
				}
				free_chunks = append(free_chunks, []int{file_size, free_space})
				file_size = -1
				free_space = -1
				file_id += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read from input puzzle inputs, RIP:", err)
	}

	free_chunks_copy := make([][]int, len(free_chunks))
	for i := range free_chunks {
		free_chunks_copy[i] = make([]int, len(free_chunks[i]))
		copy(free_chunks_copy[i], free_chunks[i])
	}

	for current_file_id := file_id - 1; current_file_id >= 0; current_file_id -= 1 {
		current_idx := 0
		file_size := file_sizes[current_file_id]
		for i, chunk := range free_chunks {
			if chunk[1] < file_size {
				current_idx += (free_chunks_copy[i][0] + free_chunks_copy[i][1])
				continue
			}
			// unset old location
			replace_idx := slices.Index(disk, current_file_id)
			replace_size := file_sizes[current_file_id]
			for j := replace_idx; j < (replace_idx + replace_size); j += 1 {
				disk[j] = -1
			}
			// replace at new location
			diff := 0
			if free_chunks_copy[i][1]-free_chunks[i][1] != 0 {
				diff = free_chunks_copy[i][1] - free_chunks[i][1]
			}
			current_idx += free_chunks_copy[i][0] + diff
			free_chunks[i][1] -= file_size
			for j := current_idx; j < (current_idx + file_size); j += 1 {
				disk[j] = current_file_id
			}
			break
		}
	}

	total := 0
	for i := 0; i < len(disk); i += 1 {
		if disk[i] != -1 {
			total += (i * disk[i])
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
