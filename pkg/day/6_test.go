package day

import "testing"

func TestSixPartOne(t *testing.T) {
	input := []string{
		"3,4,3,1,2",
	}
	result, err := SixPartOne(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 5934
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}

func TestSixPartTwo(t *testing.T) {
	input := []string{
		"3",
		"4",
		"3",
		"1",
		"2",
	}
	result, err := SixPartTwo(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 26984457539
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}
