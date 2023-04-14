package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("Error empty input.txt file")
	}
}

type pair struct {
	sensor_row, sensor_col int
	beacon_row, beacon_col int 
}

func str_int(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}

func manhattan_dist(x1, y1, x2, y2 int) int {
	xDiff := x1-x2
	yDiff := y1-y2

	return abs_int(xDiff) + abs_int(yDiff)
}

func parse_input(input string) (ans []pair) {
	// Sensor at x=2150774, y=3136587: closest beacon is at x=2561642, y=2914773
	for _, line := range strings.Split(input, "\n") {
		p := pair{}
		_, err := fmt.Sscanf(line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&p.sensor_col, &p.sensor_row, &p.beacon_col, &p.beacon_row)
		if err != nil {
			panic("parsing: " + err.Error())
		}
		ans = append(ans, p)
	}
	return ans
}

func beacon_ex_zone(input string, target_row int) {
	pairs := parse_input(input)
	blocked_coords := map[[2]int]bool{}
	for _, p := range pairs {
		man_dist := manhattan_dist(p.sensor_col, p.sensor_row, p.beacon_col, p.beacon_row)
		diff := p.sensor_row - target_row
		if diff < 0 {
			diff *= -1
		}

		blockable := man_dist - diff
		if blockable > 0 {
			for i:=0; i<= blockable; i++ {
				blocked_coords[[2]int{
					target_row, p.sensor_col - i,
				}] = true
				blocked_coords[[2]int{
					target_row, p.sensor_col + i,
				}] = true
			}
		}

	}

	for _, p := range pairs {
		delete(blocked_coords, [2]int{p.beacon_row, p.beacon_col})
	}

	fmt.Println(len(blocked_coords))
}

type parsed_sensor struct {
	sensor_row, sensor_col int
	manhattan_dist int
}

func (sensor *parsed_sensor) exclusion_range(x, y int) bool {
	return sensor.manhattan_dist >= manhattan_dist(x, y, sensor.sensor_col, sensor.sensor_row)
}

func is_reachable(sensors []parsed_sensor, c, r int) bool {
	for _, sensor :=  range sensors {
		if sensor.exclusion_range(c, r) {
			return true
		}
	}
	return false
}

func abs_int(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func beacon_ex_zone_two(input string, max_coord int) int {
	pairs := parse_input(input)
	sensors := []parsed_sensor{}
	for _, pair := range pairs {
		sensors = append(sensors, parsed_sensor{
			sensor_row: pair.sensor_row,
			sensor_col: pair.sensor_col,
			manhattan_dist: manhattan_dist(pair.sensor_col, pair.sensor_row, pair.beacon_col, pair.beacon_row),
		})
	}

	for _, sensor := range sensors {
		dist_plus_one := sensor.manhattan_dist + 1

		for r := -dist_plus_one; r <= dist_plus_one; r++ {
			target_row := sensor.sensor_row + r

			if target_row < 0 {
				continue
			} 
			if target_row > max_coord {
				break
			}

			col_offset := dist_plus_one - abs_int(r)
			col_left := sensor.sensor_col - col_offset
			col_right := sensor.sensor_col + col_offset

			if col_left >= 0 && col_left <= max_coord && !is_reachable(sensors, col_left, target_row) {
				return(col_left*max_coord + target_row)
			}
			if col_right >= 0 && col_right <= max_coord && !is_reachable(sensors, col_right, target_row) {
				return(col_right*max_coord + target_row)
			}
		}
	}
	panic("unreachable")
}


func main() {
	beacon_ex_zone(input, 2000000)
	ans := beacon_ex_zone_two(input, 4000000)
	fmt.Println(ans)
}