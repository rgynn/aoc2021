package day

import (
	"strconv"
	"strings"
	"sync"
)

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
	fishes, err := parse_fishes(input)
	if err != nil {
		return nil, err
	}
	days := 256
	for days > 0 {
		var wg sync.WaitGroup
		for x := 0; x < len(fishes); x++ {
			wg.Add(1)
			go func(x int) {
				if new := fishes[x].tick(); new != nil {
					fishes = append(fishes, new)
				}
				defer wg.Done()
			}(x)
		}
		wg.Wait()
		days--
	}
	return len(fishes), nil
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
