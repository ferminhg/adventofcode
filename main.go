package main

import (
	"bufio"
	adoc2023 "ferminhg/adventofcode/2023"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("2023/day1.input")
	//fmt.Printf("day1(input) = %d\n", adoc2023.Day1(input))
	input = readInput("2023/day2.input")
	value, power := adoc2023.Day2(input)
	fmt.Printf("day2(input) = %d, %d\n", value, power)
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
