package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) is_empty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() string {
	if s.is_empty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *Stack) peek() string {
	index := len(*s) - 1
	element := (*s)[index]
	return element
}

func (s *Stack) reverse() {
	for i := 0; i < len(*s) / 2; i++ {
		j := len(*s) - i - 1
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

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

func read_file(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(contents)
}

func split_input(lines []string) ([]string, []string, int) {
	var (
		drawings []string
		moves []string
	)

	blank_line_ind := 0
	num_stacks := 0
	for i, line := range lines {
		if string(line[0]) == " " && string(line[1]) != " " {
			num_stacks, _ = strconv.Atoi(string(line[len(line)-2]))
			blank_line_ind = i+1
			break
		}
		drawings = append(drawings, line)
	}
	for i, line := range lines {
		if i > blank_line_ind {
			moves = append(moves, line)
		}
	}

	return drawings, moves, num_stacks
}

func parse_moves(lines []string) [][]int {
	_, moves, _ := split_input(lines)
	list_moves := make([][]int, len(moves))
	for i, move := range moves {
		str := strings.ReplaceAll(move, "move ", "")
		str = strings.ReplaceAll(str, "from ", "")
		str = strings.ReplaceAll(str, "to ", "")
		move := strings.Split(str, " ")
		list_moves[i] = make([]int, 3)
		list_moves[i][0], _ = strconv.Atoi(string(move[0]))
		list_moves[i][1], _ = strconv.Atoi(string(move[1]))
		list_moves[i][2], _ = strconv.Atoi(string(move[2]))
	}
	return list_moves
}

func create_stacks(lines []string) []Stack {
	drawings, _, num_stacks := split_input(lines)
	stacks := make([]Stack, num_stacks)
	i := 0
	for _, line := range drawings {
		for j := 1; j < len(line); j += 4 {
			if string(line[j]) != " " {
				stacks[i].push(string(line[j]))
			}
			i++
		}
		i=0
	}

	fmt.Println(stacks)
	for i := range stacks {
		stacks[i].reverse()
	}
	fmt.Println(stacks)
	return stacks
}

func supply_stacks(filename string) {

	lines := read_by_line(filename)

	moves := parse_moves(lines)
	stacks := create_stacks(lines)
	
	to_return := ""

	for _, move := range moves {
		from := move[1] - 1
		to := move[2] - 1
		for i := 0; i < move[0]; i++ {
			to_move := stacks[from].pop()
			stacks[to].push(to_move)
		}
	}

	for _, stack := range stacks {
		to_return += stack.peek()
	}

	fmt.Println(to_return)
}

func supply_stacks_two(filename string) {
	file := read_file(filename)
	lines := strings.Split(file, "\n")

	moves := parse_moves(lines)
	stacks := create_stacks(lines)
	var list_to_move Stack
	to_return := ""

	for _, move := range moves {
		from := move[1] - 1
		to := move[2] - 1
		for i := 0; i < move[0]; i++ {
			list_to_move.push(stacks[from].pop())
		}
		for range list_to_move {
			stacks[to].push(list_to_move.pop())
		}
	}

	for _, stack := range stacks {
		to_return += stack.peek()
	}

	fmt.Println(to_return)
}

func main()  {
	supply_stacks("input.txt")
	supply_stacks_two("input.txt")
}
