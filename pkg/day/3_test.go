package day

import "testing"

func TestThreePartOne(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result, err := ThreePartOne(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 198
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}

func TestThreePartTwo(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result, err := ThreePartTwo(input)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := result.(int)
	if !ok {
		t.Fatal("could not cast result to int")
	}
	expected := 230
	if got != expected {
		t.Fatalf("expected answer to be: %d, got: %d", expected, got)
	}
}
