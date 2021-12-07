package day

import "testing"

func TestFivePartOne(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	result, err := FivePartOne(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 5
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}

func TestFivePartTwo(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	result, err := FivePartTwo(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 5
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}
