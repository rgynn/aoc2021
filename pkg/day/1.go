package day

import (
	"strconv"
)

func OnePartOne(input []string) (interface{}, error) {
	var last, result int
	for i, depth_input := range input {
		depth, err := strconv.Atoi(depth_input)
		if err != nil {
			return nil, err
		}
		if i != 0 && depth > last {
			result++
		}
		last = depth
	}
	return result, nil
}

func OnePartTwo(input []string) (interface{}, error) {
	var last, result int
	for i := range input {
		if len(input)-1 < i+2 {
			break
		}
		m1, err := strconv.Atoi(input[i])
		if err != nil {
			return nil, err
		}
		m2, err := strconv.Atoi(input[i+1])
		if err != nil {
			return nil, err
		}
		m3, err := strconv.Atoi(input[i+2])
		if err != nil {
			return nil, err
		}
		sum := m1 + m2 + m3
		if last < sum && i != 0 {
			result++
		}
		last = sum
	}
	return result, nil
}
