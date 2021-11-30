package part

import "fmt"

type AnswerFunc func(input []string) (interface{}, error)

func ErrorNotDefined(inputDay, inputPart int) AnswerFunc {
	return func(input []string) (interface{}, error) {
		return nil, fmt.Errorf("no function defined for day: %d, part: %d yet", inputDay, inputPart)
	}
}
