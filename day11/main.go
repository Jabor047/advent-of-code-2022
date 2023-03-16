package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func read_line_by(file string) []string {
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

type monkey struct {
	id int
	items []int
	op string
	test_num int
	true_to int
	false_to int
	items_insp int
}

func parse_input(lines []string) (ans []*monkey) {
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i <= len(lines); i+=7 {
		id, _ := strconv.Atoi(string(lines[i][7]))
		items := []int{}
		itms := strings.Split(lines[i+1][17:], ",")
		for _, item := range itms {
			item = strings.TrimSpace(item)
			itm, _ := strconv.Atoi(item)
			items = append(items, itm)
		}
		op := lines[i+2][23:]
		test_num, _ := strconv.Atoi(re.FindString(lines[i+3]))
		true_to, _ := strconv.Atoi(re.FindString(lines[i+4]))
		false_to, _ := strconv.Atoi(re.FindString(lines[i+5]))

		ans = append(ans, &monkey {
			id: id,
			items: items,
			op: op,
			test_num: test_num,
			true_to: true_to,
			false_to: false_to,
			items_insp: 0,
		})
	}

	return ans
}

func pop(a[]int) []int {
	return a[1:]
}

func mitm(file string, rounds int, manageable_worry bool) {
	lines := read_line_by(file)
	monkeys := parse_input(lines)
	re := regexp.MustCompile("[0-9]+")
	worry_level := 0
	
	big_mod := 1
	for _, m := range monkeys {
		big_mod *= m.test_num
	}

	for i:=0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				switch monkey.op[0] {
				case '*':
					num, err := strconv.Atoi(re.FindString(monkey.op))
					if err != nil {
						num = item
					}
					worry_level = item * num
				case '+':
					num, err := strconv.Atoi(re.FindString(monkey.op))
					if err != nil {
						num = item
					}
					worry_level = item + num
				}

				if manageable_worry {
					worry_level /= 3
				} else {
					worry_level %= big_mod
				}

				if worry_level % monkey.test_num == 0 {
					monkeys[monkey.true_to].items = append(monkeys[monkey.true_to].items, worry_level)
				} else {
					monkeys[monkey.false_to].items = append(monkeys[monkey.false_to].items, worry_level)
				}
				monkey.items = pop(monkey.items)
				monkey.items_insp++
			}
		}
	}

	for _, monkey := range monkeys {
		fmt.Println("Monkey", monkey.id, "inspect", monkey.items_insp, "times")
	}
	
}

func main() {
	mitm("input.txt", 20, true)
	mitm("input.txt", 10000, false)
}