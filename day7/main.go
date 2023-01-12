package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	name string
	parent_dir *dir
	child_dirs map[string]*dir
	files map[string]int
	total_size int
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

func parse_dir(cmds []string) *dir {
	root := &dir{
		name: "root",
		child_dirs: map[string]*dir{},
	}
	cur_dir := root
	c := 0

	for c < len(cmds) {
		switch cmd := cmds[c]; cmd[0:1] {
		case "$":
			if cmd == "$ ls" {
				c++
			} else {
				change_dir := strings.Split(cmd, "cd ")[1]
				change_dir = strings.TrimSpace(change_dir)
				if change_dir == ".." {
					cur_dir = cur_dir.parent_dir
				} else {
					if _, exists := cur_dir.child_dirs[change_dir]; !exists {
						cur_dir.child_dirs[change_dir] = &dir{
							name: change_dir,
							parent_dir: cur_dir,
							child_dirs: map[string]*dir{},
							files: map[string]int{},
						}
					}
					cur_dir = cur_dir.child_dirs[change_dir]
				}
				c++
			}
		default:
			if strings.HasPrefix(cmd, "dir") {
				child_dir_name := cmd[4:]
				if _, exists := cur_dir.child_dirs[child_dir_name]; !exists {
					cur_dir.child_dirs[child_dir_name] = &dir{
						name: child_dir_name,
						parent_dir: cur_dir,
						child_dirs: map[string]*dir{},
						files: map[string]int{},
					}
				}
			} else {
				parts := strings.Split(cmd, " ")
				cur_dir.files[parts[1]], _ = strconv.Atoi(parts[0])
			}
			c++
		}
	}
	calc_file_size(root)
	return root
}

func calc_file_size(cur_dir *dir) int {
	total_size := 0

	for _, child_dir := range cur_dir.child_dirs {
		total_size += calc_file_size(child_dir)
	}

	for _, size := range cur_dir.files {
		total_size += size
	}

	cur_dir.total_size = total_size

	return total_size
}

func sum_dirs_under(directory *dir, size int) int {
	sum := 0

	if directory.total_size <= size {
		sum += directory.total_size
	}

	for _, child_dir := range directory.child_dirs {
		sum += sum_dirs_under(child_dir, size)
	}

	return sum
} 

func all_dirs_over(directory *dir, size int) []int {
	var greater_than []int
	if directory.total_size >= size {
		greater_than = append(greater_than, directory.total_size)
	}

	for _, child_dir := range directory.child_dirs {
		greater_than = append(greater_than, all_dirs_over(child_dir, size)...)
	}

	return greater_than
}

func no_space(file string) {
	cmds := read_by_line(file)
	root := parse_dir(cmds)
	sum := sum_dirs_under(root, 100000)
	fmt.Println(sum)
}

func no_space_two(file string) {
	cmds := read_by_line(file)
	root := parse_dir(cmds)

	rem_space := 70000000 - root.total_size
	lim := 30000000 - rem_space
	del_size_list := all_dirs_over(root, lim)

	sort.Ints(del_size_list)
	fmt.Println(del_size_list[0])

}

func main() {
	no_space("input.txt")
	no_space_two("input.txt")
}