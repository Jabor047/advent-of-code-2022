package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calories_count() {
	// contents, err := os.ReadFile("/Users/kev_in/Projects/personal/adventofcode22/day/calories.txt")
	contents, err := os.ReadFile("./calories.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	str_list := strings.Split(string(contents), "\n\n")

	var max_list []int
	for _, i := range str_list[1:] {
		sum := 0
		i_list := strings.Split(i, "\n")

		for _, j := range i_list {
			if j != "" {
				num, err := strconv.Atoi(j)
				if err != nil {
					fmt.Println("String Conversion error", err)
				}
				sum += num
			}
		}
		max_list = append(max_list, sum)
	}
	sort.Slice(max_list, func(i, j int) bool {
		return max_list[i] > max_list[j]
	})
	top_three_total := 0
	for _, v := range max_list[:3] {
		top_three_total += v
	}
	fmt.Println(top_three_total)
}

func main(){
	calories_count()
}