package day

import (
	"errors"
	"regexp"
	"strconv"
)

func FivePartOne(input []string) (interface{}, error) {
	chart := map[int]map[int]int{}
	lines, err := parse_lines(input)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		if line.y1 == line.y2 {
			if line.x1 < line.x2 {
				for x := line.x1; x <= line.x2; x++ {
					_, ok := chart[line.y1]
					if !ok {
						chart[line.y1] = map[int]int{}
					}
					_, ok = chart[line.y1][x]
					if !ok {
						chart[line.y1][x] = 1
					} else {
						chart[line.y1][x]++
					}
				}
			} else {
				for x := line.x2; x <= line.x1; x++ {
					_, ok := chart[line.y1]
					if !ok {
						chart[line.y1] = map[int]int{}
					}
					_, ok = chart[line.y1][x]
					if !ok {
						chart[line.y1][x] = 1
					} else {
						chart[line.y1][x]++
					}
				}
			}
		}
		if line.x1 == line.x2 {
			if line.y1 < line.y2 {
				for y := line.y1; y <= line.y2; y++ {
					_, ok := chart[y]
					if !ok {
						chart[y] = map[int]int{}
					}
					_, ok = chart[y][line.x1]
					if !ok {
						chart[y][line.x1] = 1
					} else {
						chart[y][line.x1]++
					}
				}
			} else {
				for y := line.y2; y <= line.y1; y++ {
					_, ok := chart[y]
					if !ok {
						chart[y] = map[int]int{}
					}
					_, ok = chart[y][line.x1]
					if !ok {
						chart[y][line.x1] = 1
					} else {
						chart[y][line.x1]++
					}
				}
			}
		}
	}
	var result int
	for _, row := range chart {
		for _, n := range row {
			if n >= 2 {
				result++
			}
		}
	}
	return result, nil
}

func FivePartTwo(input []string) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}

type line struct {
	x1, y1, x2, y2 int
}

func parse_lines(input []string) (lines []*line, err error) {
	for _, str := range input {
		line, err := parse_line(str)
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return
}

func parse_line(input string) (*line, error) {
	rgx := regexp.MustCompile(`(\d+)`)
	lines := rgx.FindAllString(input, 4)
	x1, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, err
	}
	y1, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, err
	}
	x2, err := strconv.Atoi(lines[2])
	if err != nil {
		return nil, err
	}
	y2, err := strconv.Atoi(lines[3])
	if err != nil {
		return nil, err
	}
	return &line{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}, nil
}
