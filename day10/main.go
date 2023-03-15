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

type insts struct {
	name string
	val int
	cycles int

}

func parse_input(lines []string) (ans []insts) {
	for _, line := range lines {
		switch line[:4] {
		case "addx":
			val, _ := strconv.Atoi(line[5:])
			ans = append(ans, insts{
				name: "addx",
				val: val,
				cycles: 2,
			})
		case "noop":
			ans = append(ans, insts{
				name: "noop",
				cycles: 1,
			})
		}
	}

	return ans
}

func cathode_ray(file string) {
	lines := read_by_line(file)
	insts := parse_input(lines)

	signal_strength := 0
	value := 1
	i := 0

	for cycle := 1; cycle <= 220; cycle++ {
		if (cycle-20)%40 == 0 {
			signal_strength += value * cycle 
		}
		switch insts[i].name {
		case "addx":
			insts[i].cycles--
			if insts[i].cycles == 0 {
				value += insts[i].val
				i++
			}
		case "noop":
			i++
		}
	}

	fmt.Println(signal_strength)
}

func cathode_ray_two(file string){
	lines := read_by_line(file)
	insts := parse_input(lines)

	CRT := [6][40]string{}
	for i, rows := range CRT {
		for j := range rows {
			CRT[i][j] = "."
		}
	}
	
	cycles := 0
	for c := range insts {
		cycles += insts[c].cycles
	}

	i := 0
	value := 1
	for cycle := 1; cycle < cycles; cycle++ {
		pixel_row := (cycle - 1) / 40
		pixel_col := (cycle - 1) % 40

		sprite_left, sprite_right := value-1, value+1
		if sprite_left <= pixel_col && sprite_right >= pixel_col {
			CRT[pixel_row][pixel_col] = "#"
		}

		switch insts[i].name {
		case "addx":
			insts[i].cycles--
			if insts[i].cycles == 0 {
				value += insts[i].val
				i++
			}
		case "noop":
			i++
		}
}

	log := ""
	for _, row := range CRT {
		for _, cell := range row {
			log += cell
		}
		log += "\n"
	}

	fmt.Println(log)

}

func main(){
	cathode_ray("input.txt")
	cathode_ray_two("input.txt")
}