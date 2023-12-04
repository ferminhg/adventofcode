package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	adeventofcode2023 "ferminhg/adventofcode/2023"
)

func main() {
	input := readInput("2023/day1.input")
	fmt.Printf("day1(input) = %d\n", adeventofcode2023.Day1(input))
}

// read input from file
func readInput(filename string) []string {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
