package main

import (
	"fmt"
	"os"
	"strings"
)

func read_file(file string) string {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(contents)
}

func rock_paper_scissors(){
	scores := map[string]int {
		"X" : 1,
		"Y" : 2,
		"Z" : 3,
		"A Y": 6,
		"B Z": 6,
		"C X": 6,
		"A X": 3,
		"B Y": 3,
		"C Z": 3,
		"A Z": 0,
		"B X": 0,
		"C Y": 0,
	}

	str_list := strings.Split(read_file("./input.txt"), "\n")
	total_score := 0
	for _, i := range str_list {
		game_score := scores[i]
		play_list := strings.Split(string(i), " ")
		type_score := scores[play_list[1]]

		score := game_score + type_score
		total_score += score
	}

	fmt.Println(total_score)
}

func rock_paper_scissors_updated_rules() {
	scores := map[string]int {
		"X" : 0,
		"Y" : 3,
		"Z" : 6,
		"A" : 1,
		"B" : 2,
		"C" : 3,
		"A X": 3,
		"B X": 1,
		"C X": 2,
		"A Y": 1,
		"B Y": 2,
		"C Y": 3,
		"A Z": 2,
		"B Z": 3,
		"C Z": 1,
	}

	str_list := strings.Split(read_file("./input.txt"), "\n")
	total_score := 0
	for _, i := range str_list {
		game_score := scores[i]
		play_list := strings.Split(string(i), " ")
		type_score := scores[play_list[1]]

		score := game_score + type_score
		total_score += score
	}

	fmt.Println(total_score)
}

func main(){
	rock_paper_scissors()
	rock_paper_scissors_updated_rules()
}