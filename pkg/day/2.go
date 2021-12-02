package day

import (
	"strconv"
	"strings"
)

func TwoPartOne(input []string) (interface{}, error) {
	var x, y int
	for _, cmd := range input {
		dir, units_str := strings.Split(cmd, " ")[0], strings.Split(cmd, " ")[1]
		units, err := strconv.Atoi(units_str)
		if err != nil {
			return nil, err
		}
		switch dir {
		case "forward":
			x += units
		case "down":
			y += units
		case "up":
			y -= units
		}
	}
	return x * y, nil
}

func TwoPartTwo(input []string) (interface{}, error) {
	var x, y, aim int
	for _, cmd := range input {
		dir, units_str := strings.Split(cmd, " ")[0], strings.Split(cmd, " ")[1]
		units, err := strconv.Atoi(units_str)
		if err != nil {
			return nil, err
		}
		switch dir {
		case "forward":
			x += units
			y += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	return x * y, nil
}
