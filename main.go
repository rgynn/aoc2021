package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/rgynn/aoc2021/part"
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
	parts := map[int]map[int]part.AnswerFunc{
		0: {
			1: part.ErrorNotDefined(inputDay, inputPart),
			2: part.ErrorNotDefined(inputDay, inputPart),
		},
	}
	input, err := scan(inputDay, inputPart)
	if err != nil {
		log.Fatal(err)
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

func scan(day, part int) ([]string, error) {
	path := fmt.Sprintf("inputs/part_%d_%d", day, part)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	result := strings.Split(strings.TrimSpace(string(file)), "\n")
	if inputDebug {
		log.Println("read input:")
		for k, v := range result {
			fmt.Printf("%d\t%s\n", k, v)
		}
	}
	return result, nil
}