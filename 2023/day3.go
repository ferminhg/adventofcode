package adoc2023

import (
	"strconv"
)

type number struct {
	value    int
	row      int
	position []int
}

func newNumber(value int, row int, position []int) *number {
	return &number{value: value, row: row, position: position}
}

type checker struct {
	schematic []string
}

func newChecker(schematic []string) *checker {
	return &checker{schematic: schematic}
}

func Day3(schematic []string) int {
	numbers := make([]number, 0)
	for i := 0; i < len(schematic); i++ {
		numbers = append(numbers, extractNumbersByRow(schematic[i], i)...)
	}
	checker := newChecker(schematic)
	return checker.getSum(numbers)
}

type pos struct {
	x int
	y int
}

func Day3PartB(schematic []string) int64 {
	numbers := make([]number, 0)
	specialSymbolsPositions := make([]pos, 0)
	for i := 0; i < len(schematic); i++ {
		//fmt.Printf("numbers[%d] = %v\n", i, schematic[i])
		numbers = append(numbers, extractNumbersByRow(schematic[i], i)...)
		specialSymbolsPositions = append(specialSymbolsPositions, extractSpecialSymbolsPositions(schematic[i], i)...)
	}
	checker := newChecker(schematic)
	return checker.getSumPartB(numbers, specialSymbolsPositions)
}

func extractSpecialSymbolsPositions(s string, x int) []pos {
	positions := make([]pos, 0)
	for y, char := range s {
		if char == '*' {
			positions = append(positions, pos{x, y})
		}
	}
	return positions
}

func extractNumbersByRow(s string, row int) []number {
	numbers := make([]number, 0)
	number := ""
	positions := make([]int, 0)

	for j, char := range s {
		if isDigit(char) {
			number += string(char)
			positions = append(positions, j)
		} else {
			if number != "" {
				numberInt, _ := strconv.Atoi(number)
				numbers = append(numbers, *newNumber(numberInt, row, positions))
			}
			number = ""
			positions = make([]int, 0)
		}
	}

	if number != "" {
		numberInt, _ := strconv.Atoi(number)
		numbers = append(numbers, *newNumber(numberInt, row, positions))
	}

	return numbers
}

func (c *checker) getSum(numbers []number) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		if c.isValidNumber(numbers[i]) {
			//fmt.Printf("number[%d] = %v\n", i, numbers[i])
			sum += numbers[i].value
		}
	}
	return sum
}

func (c *checker) getSumPartB(numbers []number, specialSymbolsPositions []pos) int64 {
	sum := int64(0)
	//fmt.Printf("specialSymbolsPositions = %v\n", specialSymbolsPositions)
	for _, p := range specialSymbolsPositions {
		value := calculateValueByPosition(numbers, p)
		sum += value
	}
	return sum
}

func calculateValueByPosition(numbers []number, p pos) int64 {
	value := int64(0)
	for _, n := range numbers {
		for _, v := range n.position {
			// la posicion de n es adyacente a la posicion de n
			//if (n.row == p.x-1 || n.row == p.x || n.row == p.x+1) && (v == p.y-1 || v == p.y || v == p.y+1) {
			//	value += int64(n.value)
			//}
			// la posicion de n es la misma arriba o bajo de la posicion de n
			// la posicion es adyacende a la posicion de n en la misma fila, en la de arriba y la de abajo
			if (n.row == p.x-1 || n.row == p.x || n.row == p.x+1) && (v == p.y) {
				value += int64(n.value)
			}
		}
	}
	return value
}
func (c *checker) isValidNumber(n number) bool {
	return c.reviewRows(n, n.row) ||
		c.reviewRows(n, n.row-1) ||
		c.reviewRows(n, n.row+1)
}

func (c *checker) reviewRows(n number, row int) bool {
	for i, v := range n.position {
		if i == 0 && c.checkPositionIsSymbol(row, v-1) {
			return true
		}
		if c.checkPositionIsSymbol(row, v) {
			return true
		}
		if i == len(n.position)-1 && c.checkPositionIsSymbol(row, v+1) {
			return true
		}
	}
	return false
}

func (c *checker) checkPositionIsSymbol(x int, y int) bool {
	if x < 0 || y < 0 || x >= len(c.schematic) || y >= len(c.schematic[x]) {
		return false
	}
	return isSymbol(rune(c.schematic[x][y]))
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isDot(s rune) bool {
	return s == '.'
}

func isSymbol(s rune) bool {
	return !isDigit(s) && !isDot(s)
}
