package part

import "testing"

func TestDayZeroPartOne(t *testing.T) {
	input := []string{}
	result, err := DayZeroPartOne(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Fatal(err)
	}
}

func TestDayZeroPartTwo(t *testing.T) {
	input := []string{}
	result, err := DayZeroPartTwo(input)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Fatal(err)
	}
}
