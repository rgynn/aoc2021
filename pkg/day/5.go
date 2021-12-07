package day

import (
	"regexp"
	"strconv"
)

func FivePartOne(input []string) (interface{}, error) {
	lines, err := parse_lines(input)
	if err != nil {
		return nil, err
	}
	chart := [1000][1000]int{}
	chart = chart_hv(lines, chart)
	return calc_score_five(chart), nil
}

func FivePartTwo(input []string) (interface{}, error) {
	lines, err := parse_lines(input)
	if err != nil {
		return nil, err
	}
	chart := [1000][1000]int{}
	chart = chart_hv(lines, chart)
	chart = chart_diag(lines, chart)
	return calc_score_five(chart), nil
}

type line struct {
	x1, y1, x2, y2 int
}

func chart_diag(lines []*line, chart [1000][1000]int) [1000][1000]int {
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			continue
		}
		x, y := line.x1, line.y1
		if line.x1 > line.x2 {
			if line.y1 < line.y2 {
				for x >= line.x2 && y <= line.y2 {
					chart[y][x]++
					x--
					y++
				}
			} else {
				for x >= line.x2 && y >= line.y2 {
					chart[y][x]++
					x--
					y--
				}
			}
		} else {
			if line.y1 < line.y2 {
				for x <= line.x2 && y <= line.y2 {
					chart[y][x]++
					x++
					y++
				}
			} else {
				for x <= line.x2 && y >= line.y2 {
					chart[y][x]++
					x++
					y--
				}
			}
		}
	}
	return chart
}

func chart_hv(lines []*line, chart [1000][1000]int) [1000][1000]int {
	for _, line := range lines {
		if line.y1 == line.y2 {
			if line.x1 < line.x2 {
				for x := line.x1; x <= line.x2; x++ {
					chart[line.y1][x]++
				}
			} else {
				for x := line.x2; x <= line.x1; x++ {
					chart[line.y1][x]++
				}
			}
		}
		if line.x1 == line.x2 {
			if line.y1 < line.y2 {
				for y := line.y1; y <= line.y2; y++ {
					chart[y][line.x1]++
				}
			} else {
				for y := line.y2; y <= line.y1; y++ {
					chart[y][line.x1]++
				}
			}
		}
	}
	return chart
}

func calc_score_five(chart [1000][1000]int) (result int) {
	for _, row := range chart {
		for _, n := range row {
			if n >= 2 {
				result++
			}
		}
	}
	return
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
