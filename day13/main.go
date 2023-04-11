package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func parse_input(input string) (ans [][]string) {
	for _, pair := range strings.Split(input, "\n\n") {
		pair_list := strings.Split(pair, "\n")
		ans = append(ans, pair_list)
	}
	return ans
}

func unmarshal(left, right string) (any, any) {
	var l, r any

	if err := json.Unmarshal([]byte(left), &l); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(right), &r); err != nil {
		log.Fatal(err)
	}

	return l, r
}

func compare(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		l = []any{left}
	case !rok:
		r = []any{right}
	}

	for i := 0; i < len(r) && i < len(l); i++ {
		if res := compare(l[i], r[i]); res != 0 {
			return res
		}
	}

	return len(l) - len(r)
}

func distress_signal(input string) {
	ans := parse_input(input)
	sum := 0
	for idx, pair := range ans {
		var left, right = unmarshal(pair[0], pair[1])
		if compare(left, right) <= 0 {
			sum += idx+1
		}
	}
	fmt.Println(sum)
}

func distress_signal_two(input string) {
	ans := parse_input(input)
	key := 1
	var list []any

	for _, pair := range ans {
		l, r := unmarshal(pair[0], pair[1])
		list = append(list, l, r)
	}
	divider_pkts := []any{[]any{2.}}
	divider_pkts_two := []any{[]any{6.}}

	list = append(list, divider_pkts, divider_pkts_two)
	sort.Slice(list, func(i, j int) bool {
		return compare(list[i], list[j]) < 0
	})
	for idx, lst := range list {
		if fmt.Sprint(lst) == "[[2]]" || fmt.Sprint(lst) == "[[6]]" {
			key *= idx+1
		}
	}
	fmt.Println(key)
}

func main() {
	distress_signal(input)
	distress_signal_two(input)
}