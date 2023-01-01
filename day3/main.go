package main

import (
	"fmt"
	"os"
	"strings"
)

func read_file(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(contents)
}

func char_to_int(char rune) int {
	int_val := int(char)
	if int_val > 90 {
		return int_val - 96
	} else {
		return int_val - 38
	}
}

func bag_rego(filename string) {
	file := read_file(filename)
	str_list := strings.Split(strings.TrimSpace(file), "\n")
	sum := 0 
	for _, content := range str_list {
		length := len(content)
		first_half := content[:length/2]
		seconf_half := content[length/2:]

		for _, c := range first_half {
			if strings.Contains(seconf_half, string(c)) {
				sum += char_to_int(c)
				break
			}
		}
	}
	fmt.Println(sum)
}

func bag_rego_two(filename string) {
	file := read_file(filename)
	str_list := strings.Split(strings.TrimSpace(file), "\n")
	m := make(map[rune][3]int)
	sum := 0
	pointer := 0
	for _, content := range str_list {
		for _, c := range content {
			if val, exists := m[c]; exists {
				val[pointer] = 1
				m[c] = val
				if val[0] == 1 && val[1] == 1 && val[2] == 1 {
					sum += char_to_int(c)
					break
				}
			} else {
				m[c] = [3]int{0, 0, 0}
				val := m[c]
				val[pointer] = 1
				m[c] = val
			}
		}

		pointer += 1
		if pointer == 3 {
			pointer = 0
			m = make(map[rune][3]int)
		}
	}
	fmt.Println(sum)
}

func main() {
	bag_rego("input.txt")
	bag_rego_two("input.txt")
}