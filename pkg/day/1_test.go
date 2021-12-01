package day

import "testing"

func TestOnePartOne(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	result, err := OnePartOne(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 7
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}

func TestOnePartTwo(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	result, err := OnePartTwo(input)
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
