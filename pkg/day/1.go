package day

import (
	"errors"
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
	return nil, errors.New("not implemented yet")
}
