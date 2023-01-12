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

type inst struct {
	dir string
	steps int
}

func parse_input(lines []string) (o_p []inst) {
	for _, line := range lines {
		val, _ := strconv.Atoi(line[2:])
		o_p = append(o_p, inst{
			dir: line[:1],
			steps: val,
		})
	}

	return o_p
}

type node struct {
	coords [2]int
	next *node
}

type rope struct {
	head, tail *node
}

func init_rope(length int) rope {
	head := &node{}
	itr := head
	
	for i := 1; i < length; i++ {
		itr.next = &node{}
		itr = itr.next
	}

	return rope{
		head: head,
		tail: itr,
	}
}

func abs_int(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func (n *node) update_next(){
	if n.next == nil {
		return
	}

	row_diff := n.coords[0] - n.next.coords[0]
	col_diff := n.coords[1] - n.next.coords[1]

	// if either row or col diff is > 1, then that dimension HAS to move
	// additionally, if the other diff is not zero, it needs to be
	// adjusted to move diagonally

	if abs_int(row_diff) > 1 && abs_int(col_diff) > 1 {
		/*  0 1 2
			H . T
			diff = head - tail = -2
			want to move tail (2) to (1), so add diff / 2
			T . H
			diff = 2 - 0 = 2
			tail (0) + 2/2 = 1, checks out still
		*/
		n.next.coords[0] += row_diff / 2
		n.next.coords[1] += col_diff / 2
	} else if abs_int(row_diff) > 1 {
		n.next.coords[0] += row_diff / 2
		n.next.coords[1] += col_diff
	} else if abs_int(col_diff) > 1 {
		n.next.coords[0] += row_diff
		n.next.coords[1] += col_diff / 2
	} else {
		return
	}

	n.next.update_next()
}

func (r rope) move_one_space(dir string) {
	diffs := map[string][2]int{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}
	diff := diffs[dir]
	r.head.coords[0] += diff[0]
	r.head.coords[1] += diff[1]

	r.head.update_next()
}

func rope_bridge(file string, rope_len int) {
	lines := read_by_line(file)
	insts := parse_input(lines)

	rope := init_rope(rope_len)

	visited := map[[2]int]bool{}
	for _, inst := range insts {
		for inst.steps > 0 {
			rope.move_one_space(inst.dir)

			visited[rope.tail.coords] = true

			inst.steps--
		}
	}
	fmt.Println(len(visited))
}

func main() {
	rope_bridge("input.txt", 2)
	rope_bridge("input.txt", 10)
}