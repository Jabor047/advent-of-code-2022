package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func max_int(nums ...int) int {
	num := nums[0]
	for _, v := range nums {
		if v > num {
			num = v
		}
	}
	return num
}

func min_int(nums ...int) int {
	num := nums[0]
	for _, v := range nums {
		if v < num {
			num = v
		}
	}
	return num
}

func parse_input(input string)(matrix [][]string, init_col int) {
	coord_sets := [][][2]int{}
	lowest_col := math.MaxInt64
	highest_row := 0
	for _, line := range strings.Split(input, "\n") {
		input_coords := strings.Split(line, " -> ")
		coords := [][2]int{}
		for _, input_coord := range input_coords {
			str_coords := strings.Split(input_coord, ",")
			col, _ := strconv.Atoi(str_coords[0])
			row, _ := strconv.Atoi(str_coords[1])
			coord := [2]int{
				col, row,
			}

			coords = append(coords, coord)
			lowest_col = min_int(lowest_col, col)
			highest_row = max_int(highest_row, row)
		}
		coord_sets = append(coord_sets, coords)
	}

	extra_space := 150

	highest_col := 0
	for s, set := range coord_sets {
		for i := range set {
			coord_sets[s][i][0] -= lowest_col - extra_space
			highest_col = max_int(highest_col, coord_sets[s][i][0])
		}
	}

	matrix = make([][]string, highest_row+3)
	for r := range matrix {
		matrix[r] = make([]string, highest_col+extra_space*2)
	}

	for _, set := range coord_sets {
		for i := 1; i < len(set); i++ {
			cols := []int{set[i-1][0], set[i][0]}
			rows := []int{set[i-1][1], set[i][1]}

			sort.Ints(cols)
			sort.Ints(rows)

			if cols[0] == cols[1] {
				for r := rows[0];r <= rows[1];r++ {
					matrix[r][cols[0]] = "#"
				}
			} else if rows[0] == rows[1] {
				for c := cols[0]; c <= cols[1]; c++ {
					matrix[rows[0]][c] = "#"
				}
			}
		}
	}

	init_col = 500 - lowest_col + extra_space
	matrix[0][init_col] = "+"

	for i, r := range matrix {
		for j := range r {
			if matrix[i][j] == "" {
				matrix[i][j] = "."
			}
		}
	}

	// printMatrix(matrix)
	return matrix, init_col
}

func printMatrix(matrix [][]string) {
	for _, r := range matrix {
		fmt.Println(r)
	}
}

func drop_sand(matrix [][]string, init_col int) (abyss bool) {
	r, c := 0, init_col

	for r < len(matrix)-1 {
		below := matrix[r+1][c]
		diag_left := matrix[r+1][c-1]
		diag_right := matrix[r+1][c+1]

		if below == "." {
			r++
		} else if diag_left == "." {
			r++
			c--
		} else if diag_right == "." {
			r++
			c++
		} else {
			matrix[r][c] = "o"
			return false
		}
	}

	printMatrix(matrix)
	return true
}

func regolith(input string) {
	mat, init_col :=  parse_input(input)
	count := 0

	for !drop_sand(mat, init_col) {
		count++
	}
	fmt.Println(count)
}

func regolith_two(input string) {
	mat, init_col :=  parse_input(input)
	count := 0

	for i := range mat[0] {
		mat[len(mat)-1][i] = "#"
	}

	for !drop_sand(mat, init_col) {
		count++
		if mat[0][init_col] == "o" {
			break
		}
	}
	fmt.Println(count)
}

func main() {
	regolith(input)
	regolith_two(input)
}