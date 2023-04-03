package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	// "github.com/alexchao26/advent-of-code-go/cast"
)

func read_by_line(file string) []string {
	lines := []string{}
	contents, err := os.Open(file)
	if err != nil {
		fmt.Println("Error reading file", err)
		return lines
	}

	defer contents.Close()

	scanner := bufio.NewScanner(contents)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}


func parse_input(lines []string) (ans [][]string) {
	for _, line := range lines {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}


func diff_b2n_letters(x, y string) int {
	if x == "S" {
		x = "a"
	}
	if y == "S" {
		y = "a"
	}
	if y == "E" {
		y = "z"
	}
	if x == "E" {
		x = "z"
	}

	return int(y[0]) - int(x[0])
}

var diffs = [4][2] int {
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type character struct {
	row_ind int
	col_ind int
	steps int
}


func hill_climb(file string) int {
	lines := read_by_line(file)
	chars := parse_input(lines)
	
	var queue []*character

start:
	for r, rows := range chars {
		for c, char := range rows {
			if char == "S" {
				queue = append(queue, &character{
					row_ind: r,
					col_ind: c,
					steps: 0,
				})
				break start
			}
		}
	}
	seen := map[[2]int]bool{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
			
		if seen[[2]int{front.row_ind, front.col_ind}] {
			continue
		}
	
		seen[[2]int{front.row_ind, front.col_ind}] = true
	
		if chars[front.row_ind][front.col_ind] == "E" {
			return front.steps
		}

		for _, d := range diffs {
			next_r, next_c := front.row_ind+d[0], front.col_ind+d[1]
			if next_r >= 0 && next_r < len(chars) &&  next_c >= 0 && next_c < len(chars[0]) {
				letter_diff := diff_b2n_letters(chars[front.row_ind][front.col_ind], chars[next_r][next_c])
				if letter_diff <= 1 {
					queue = append(queue, &character{
						row_ind: next_r,
						col_ind: next_c,
						steps: front.steps+1})
				}
			}
		}
		
	}
	return -1
}

func hill_climb_two(file string) int {
	lines := read_by_line(file)
	chars := parse_input(lines)
	
	var queue []*character

end:
	for r, rows := range chars {
		for c, char := range rows {
			if char == "E" {
				queue = append(queue, &character{
					row_ind: r,
					col_ind: c,
					steps: 0,
				})
				break end
			}
		}
	}

	seen := map[[2]int]bool{}

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
	
		if seen[[2]int{front.row_ind, front.col_ind}] {
			continue
		}

		seen[[2]int{front.row_ind, front.col_ind}] = true

		if chars[front.row_ind][front.col_ind] == "a" {
			return front.steps
		}

		for _, d := range diffs {
			next_r, next_c := front.row_ind + d[0], front.col_ind + d[1]
			if next_r >= 0 && next_r < len(chars) && next_c >= 0 && next_c < len(chars[0]) {
				letter_diff := diff_b2n_letters(chars[front.row_ind][front.col_ind], chars[next_r][next_c])
				
				if letter_diff >= -1 {
					queue = append(queue, &character{
						row_ind: next_r,
						col_ind: next_c,
						steps: front.steps + 1,
					})
				}
			}
		}

	}
	
	return -1
}

func main() {
	fmt.Println(hill_climb("input.txt"))
	fmt.Println(hill_climb_two("input.txt"))
}