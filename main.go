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
	//value, power := adoc2023.Day2(input)
	input = readInput("2023/day3.input")
	//value := adoc2023.Day3(input)
	//fmt.Printf("day3(input) = %d\n", value)
	//value := adoc2023.Day3(input)
	//fmt.Printf("day3(input) = %d\n", value)
	input = readInput("2023/day4.input")
	value := adoc2023.Day4Part1(input)
	fmt.Printf("day4(input) = %d\n", value)
	value = adoc2023.Day4Part2(input)
	fmt.Printf("day4partB(input) = %d\n", value)
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
