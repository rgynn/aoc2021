package day

import (
	"errors"
	"strconv"
	"strings"
)

// After one day, its internal timer would become 2.
// After another day, its internal timer would become 1.
// After another day, its internal timer would become 0.
// After another day, its internal timer would reset to 6, and it would create a new lanternfish with an internal timer of 8.
// After another day, the first lanternfish would have an internal timer of 5, and the second lanternfish would have an internal timer of 7.

type fish struct {
	offset int
}

func (f *fish) tick() *fish {
	if f.offset == 0 {
		f.offset = 6
		return &fish{offset: 8}
	} else {
		f.offset--
	}
	return nil
}

func SixPartOne(input []string) (interface{}, error) {
	fishes, err := parse_fishes(input)
	if err != nil {
		return nil, err
	}
	days := 80
	for days > 0 {
		for i := range fishes {
			if new := fishes[i].tick(); new != nil {
				fishes = append(fishes, new)
			}
		}
		days--
	}
	return len(fishes), nil
}

func SixPartTwo(input []string) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}

func parse_fishes(input []string) (result []*fish, err error) {
	for _, f := range strings.Split(input[0], ",") {
		offset, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		result = append(result, &fish{offset: offset})
	}
	return
}
