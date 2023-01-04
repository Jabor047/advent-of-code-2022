package main

import (
	"os"
	"fmt"
)

func read_file(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(contents)
}

func tuning_trouble(file string, sub_size int) {
	lines := read_file(file)
	ind_map := make(map[string] int) 

	start := 0
	end := 0

	ind_map[string(lines[end])] = 0

	for end < len(lines) {
		if end - start + 1 == sub_size {
			break
		}
		end++
		current_char := string(lines[end])
		if _, exists := ind_map[current_char]; exists {
			new_start := ind_map[current_char] + 1
			for start < new_start {
				ind_map[string(lines[start])] = -1
				start++
			}
			ind_map[current_char] = end
		} else {
			ind_map[current_char] = end
		}
	}
	fmt.Println(ind_map)
	fmt.Println(end + 1)
}

func main() {
	tuning_trouble("input.txt", 4)
	tuning_trouble("input.txt", 14)
}

