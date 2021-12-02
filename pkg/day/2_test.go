package day

import "testing"

func TestTwoPartOne(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	result, err := TwoPartOne(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 150
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}

func TestTwoPartTwo(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	result, err := TwoPartTwo(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 900
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}
