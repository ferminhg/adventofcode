package adeventofcode2023

import (
	"strconv"
	"strings"
)

func Day1(input []string) int {
	total := 0
	for _, value := range input {
		//fmt.Printf("value: %s\n", value)
		calibrationValue, _ := extractCalibration(value)
		//fmt.Printf("value: %d\n", calibrationValue)
		total += calibrationValue
	}
	return total
}

func extractCalibration(line string) (int, error) {
	firstDigit, lastDigit := rune(-1), rune(-1)
	for i, char := range line {
		if char >= '1' && char <= '9' {
			if firstDigit == -1 {
				firstDigit = char
			}
			lastDigit = char
		}
		if prefixLen := startWithAny(line[i:]); prefixLen > 0 {
			if firstDigit == -1 {
				firstDigit = rune('0') + rune(prefixLen)
			}
			lastDigit = rune('0') + rune(prefixLen)
		}
	}
	concatenatedDigits := string(firstDigit) + string(lastDigit)
	result, err := strconv.Atoi(concatenatedDigits)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func startWithAny(line string) int {
	prefixes := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for num, prefix := range prefixes {
		if strings.HasPrefix(line, prefix) {
			return num + 1
		}
	}
	return 0
}
