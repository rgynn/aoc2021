package part

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type AnswerFunc func(input []string) (interface{}, error)

func ErrorNotDefined(inputDay, inputPart int) AnswerFunc {
	return func(input []string) (interface{}, error) {
		return nil, fmt.Errorf("no function defined for day: %d, part: %d yet", inputDay, inputPart)
	}
}

func Scan(debug, test bool, day, part int) ([]string, error) {
	var prefix string
	if test {
		prefix = "../../"
	}
	path := fmt.Sprintf("%sinputs/day_%d_%d", prefix, day, part)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	result := strings.Split(strings.TrimSpace(string(file)), "\n")
	if debug {
		log.Println("read input:")
		for k, v := range result {
			fmt.Printf("%d\t%s\n", k, v)
		}
	}
	return result, nil
}
