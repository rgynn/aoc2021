package day

import (
	"strconv"
	"strings"
)

func SixPartOne(input []string) (interface{}, error) {
	return calc_fishes(80, input)
}

func SixPartTwo(input []string) (interface{}, error) {
	return calc_fishes(256, input)
}

func calc_fishes(num_days int, input []string) (int, error) {
	days, err := parse_fish_days(input)
	if err != nil {
		return 0, err
	}
	for x := 0; x < num_days; x++ {
		days = tick_fish_day(days)
	}
	var sum int
	for _, day := range days {
		sum += day
	}
	return sum, nil
}

func tick_fish_day(days []int) []int {
	new_days := make([]int, 9)
	for x := 1; x < 9; x++ {
		new_days[x-1] = days[x]
	}
	new_days[6] += days[0]
	new_days[8] += days[0]
	return new_days
}

func parse_fish_days(input []string) ([]int, error) {
	days := make([]int, 9)
	for _, f := range strings.Split(input[0], ",") {
		num, err := strconv.Atoi(f)
		if err != nil {
			return days, err
		}
		days[num]++
	}
	return days, nil
}
