package adeventofcode2023

import "testing"

func TestDay1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := 142
	actual := Day1(input)
	if actual != expected {
		t.Errorf("Day1(input) = %d, want %d", actual, expected)
	}
}

func TestDay1Part2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281
	actual := Day1(input)
	if actual != expected {
		t.Errorf("Day1(input) = %d, want %d", actual, expected)
	}
}
