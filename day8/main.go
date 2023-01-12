package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_by_line(file string) []string {
	lines := []string{}
	contents, err := os.Open(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return lines
	}

	defer contents.Close()

	scanner := bufio.NewScanner(contents)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parse_input(lines []string) ([][]int) {
	var matrix [][]int
	for _, line := range lines {
		var line_arr []int
		for _, j := range line {
			num, _ := strconv.Atoi(string(j)) 
			line_arr = append(line_arr, num)
		}
		matrix = append(matrix, line_arr)
	}
	
	return matrix
}

func tree_house(file string) {
	lines := read_by_line(file)
	matrix := parse_input(lines)

	visible_coords := map[[2]int]string{}

	for row := 1; row < len(matrix) - 1; row++ {
		highest_from_left := -1
		for col := 0; col < len(matrix[0]) - 1; col++ {
			height := matrix[row][col]
			if height > highest_from_left {
				highest_from_left = height
				visible_coords[[2]int{row, col}] = "L"
			}
		}

		highest_from_right := -1 
		for col := len(matrix[0]) - 1; col > 0; col-- {
			height := matrix[row][col]
			if height > highest_from_right {
				highest_from_right = height
				visible_coords[[2]int{row, col}] = "R"
			}
		}
	}

	for col := 1; col <  len(matrix[0]) - 1; col++ {
		highest_from_bottom := -1
		for row := len(matrix) - 1; row > 0; row-- {
			height := matrix[row][col]
			if height > highest_from_bottom {
				highest_from_bottom = height
				visible_coords[[2]int{row, col}] = "B"
			}
		}

		highest_from_top := -1
		for row := 0; row < len(matrix); row++ {
			height := matrix[row][col]
			if height > highest_from_top {
				highest_from_top = height
				visible_coords[[2]int{row, col}] = "T"
			}
		}
	}

	fmt.Println(len(visible_coords)+4)
}

func visible(mat [][]int, r, c, dr, dc int) int {
	count := 0
	start_height := mat[r][c]
	r += dr
	c += dc
	for r >= 0 && r < len(mat) && c >= 0 && c < len(mat[0]) {
		height := mat[r][c]
		if height < start_height {
			count ++
		} else {
			count++
			break
		}

		r += dr
		c += dc
	}

	return count
}

func tree_house_scenic(file string) {
	lines := read_by_line(file)
	matrix := parse_input(lines)

	best_score := 0
	for r := 1; r < len(matrix) - 1; r++ {
		for c := 1; c < len(matrix) - 1; c++ {
			score := visible(matrix, r, c, -1, 0)
			score *= visible(matrix, r, c, 1, 0)
			score *= visible(matrix, r, c, 0, -1)
			score *= visible(matrix, r, c, 0, 1)


			if score > best_score {
				best_score = score
			}
		}
	}

	fmt.Println(best_score)
}

func main() {
	tree_house("input.txt")
	tree_house_scenic("input.txt")
}