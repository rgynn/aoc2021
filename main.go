package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rgynn/aoc2021/pkg/day"
	"github.com/rgynn/aoc2021/pkg/part"
)

var (
	inputDay   int
	inputPart  int
	inputDebug bool
)

func init() {
	flag.IntVar(&inputDay, "d", 0, "day")
	flag.IntVar(&inputPart, "p", 0, "part")
	flag.BoolVar(&inputDebug, "debug", false, "debug mode")
	flag.Parse()
}

func main() {
	input, err := part.Scan(inputDebug, false, inputDay, inputPart)
	if err != nil {
		log.Fatal(err)
	}
	parts := map[int]map[int]part.AnswerFunc{
		1: {
			1: day.OnePartOne,
			2: day.OnePartTwo,
		},
		2: {
			1: day.TwoPartOne,
			2: day.TwoPartTwo,
		},
		3: {
			1: day.ThreePartOne,
			2: day.ThreePartTwo,
		},
		4: {
			1: day.FourPartOne,
			2: day.FourPartTwo,
		},
	}
	answer_func, ok := parts[inputDay][inputPart]
	if !ok {
		answer_func = part.ErrorNotDefined(inputDay, inputPart)
	}
	answer, err := answer_func(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day: %d, part: %d, answer: %v\n", inputDay, inputPart, answer)
}
