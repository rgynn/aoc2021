package day

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FourPartOne(input []string) (interface{}, error) {
	cmds := strings.Split(input[0], ",")
	input = clean_input(input)
	boards := parse_boards(input)
	for _, cmd := range cmds {
		for i := range boards {
			dutta(cmd, i, boards)
			if bingo := check_bingo(i, boards); bingo {
				return calc_score(i, boards, cmd)
			}
		}
	}
	return 0, errors.New("no winner found")
}

func FourPartTwo(input []string) (interface{}, error) {
	cmds := strings.Split(input[0], ",")
	input = clean_input(input)
	boards := parse_boards(input)
	winners := []int{}
	for _, cmd := range cmds {
		for i := range boards {
			dutta(cmd, i, boards)
			if bingo := check_bingo(i, boards); bingo {
				if !includes_board(i, winners) {
					winners = append(winners, i)
				}
				if len(winners) == len(boards) {
					return calc_score(i, boards, cmd)
				}
			}
		}
	}
	return 0, errors.New("no winner found")
}

func includes_board(i int, winners []int) bool {
	for _, winner := range winners {
		if i == winner {
			return true
		}
	}
	return false
}

func calc_score(i int, boards [][][]string, cmd string) (int, error) {
	score, err := calc_score_for_board(boards[i])
	if err != nil {
		return 0, err
	}
	win_cmd, err := strconv.Atoi(cmd)
	if err != nil {
		return 0, err
	}
	return score * win_cmd, nil
}

func calc_score_for_board(board [][]string) (result int, err error) {
	for _, row := range board {
		for _, num := range row {
			if strings.HasPrefix(num, "x") {
				continue
			}
			digit, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			result += digit
		}
	}
	return result, nil
}

func check_bingo(i int, boards [][][]string) bool {
	return check_bingo_h(i, boards) ||
		check_bingo_v(i, boards) ||
		false
}

func check_bingo_h(i int, boards [][][]string) bool {
rowloop:
	for _, row := range boards[i] {
		for _, num := range row {
			if !strings.HasPrefix(num, "x") {
				continue rowloop
			}
		}
		return true
	}
	return false
}

func check_bingo_v(i int, boards [][][]string) bool {
	for x := 0; x < 5; x++ {
		if strings.HasPrefix(boards[i][0][x], "x") &&
			strings.HasPrefix(boards[i][1][x], "x") &&
			strings.HasPrefix(boards[i][2][x], "x") &&
			strings.HasPrefix(boards[i][3][x], "x") &&
			strings.HasPrefix(boards[i][4][x], "x") {
			return true
		}
	}
	return false
}

func dutta(cmd string, i int, boards [][][]string) {
	for y := range boards[i] {
		for x := range boards[i][y] {
			if boards[i][y][x] == cmd {
				boards[i][y][x] = fmt.Sprintf("x%s", boards[i][y][x])
				return
			}
		}
	}
}

func clean_input(input []string) (result []string) {
	for _, line := range input[1:] {
		if line != "" {
			result = append(result, line)
		}
	}
	return
}

func parse_boards(input []string) (boards [][][]string) {
	board_size := 5
	for x := 0; x < len(input)/board_size; x++ {
		start := (x * board_size)
		end := (x * board_size) + board_size
		boards = append(boards, parse_board(input[start:end]))
	}
	return boards
}

func parse_board(input []string) (board [][]string) {
	rgx := regexp.MustCompile(`(\d+)`)
	for _, line := range input {
		board = append(board, rgx.FindAllString(line, 5))
	}
	return
}
