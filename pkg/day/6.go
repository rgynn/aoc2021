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

func calc_fishes(days int, input []string) (int, error) {
	fishes, err := parse_fishes(input)
	if err != nil {
		return 0, err
	}
	for days > 0 {
		for x := range fishes {
			switch fishes[x] {
			case 0:
				fishes[x] = 6
				fishes = append(fishes, 8)
			default:
				fishes[x]--
			}
		}
		days--
	}
	return len(fishes), nil
}

func parse_fishes(input []string) (result []uint8, err error) {
	for _, f := range strings.Split(input[0], ",") {
		offset, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		result = append(result, uint8(offset))
	}
	return
}
