package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func read_file(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(contents)
}

func parse_str_list(str_list []string) (lim_list [][]int) {
	for _, line := range str_list {
		sections := strings.Split(line, ",")
	
		sec_one := strings.Split(sections[0], "-")
		sec_two := strings.Split(sections[1], "-")

		sec_one_l, _ := strconv.Atoi(sec_one[0])
		sec_one_u, _ := strconv.Atoi(sec_one[1])
		sec_two_l, _ := strconv.Atoi(sec_two[0])
		sec_two_u, _ := strconv.Atoi(sec_two[1])

		lim_list = append(lim_list, []int{
			sec_one_l, sec_one_u,
			sec_two_l, sec_two_u,
		})
	}

	return lim_list
}

func does_pair1_con_pair2(pair1, pair2 []int) bool {
	return pair1[0] >= pair2[0] && pair1[1] <= pair2[1]
}

func does_overlap(pair1, pair2 []int) bool {
	if pair1[0] > pair2[0] {
		pair1, pair2 = pair2, pair1
	}
	return pair2[0] <= pair1[1]
}

func camp_cleanup(filename string){
	file := read_file(filename)
	str_list := strings.Split(file, "\n")
	var pairs_count int = 0
	lim_list := parse_str_list(str_list)

	for _, lim := range lim_list {
		if does_pair1_con_pair2(lim[:2], lim[2:]) || does_pair1_con_pair2(lim[2:], lim[:2]) {
			pairs_count++
		}
	}
	fmt.Println(pairs_count)
}

func camp_cleanup_two(filename string) {
	file := read_file(filename)
	str_list := strings.Split(file, "\n")
	var overlap_count int = 0
	lim_list := parse_str_list(str_list)

	for _, lim := range lim_list {
		if does_overlap(lim[:2], lim[2:]) {
			overlap_count++
		}
	}
	fmt.Println(overlap_count)
}

func main() {
	camp_cleanup("input.txt")
	camp_cleanup_two("input.txt")
}
