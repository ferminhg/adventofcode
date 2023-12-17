package adoc2023

import (
	"strconv"
	"strings"
)

func Day4Part1(input []string) int64 {
	result := int64(0)
	for _, rawInput := range input {
		rawNumbers := strings.Split(rawInput, ":")
		numbers := strings.Split(rawNumbers[1], "|")
		winningNumber := convertStringArrayToIntArray(strings.Split(numbers[0], " "))
		numbersYouHave := convertStringArrayToIntArray(strings.Split(numbers[1], " "))

		result += checkPrice(winningNumber, numbersYouHave)
		//fmt.Println(winningNumber, numbersYouHave)
	}
	return result
}

func Day4Part2(input []string) int64 {
	result := int64(0)
	maxCard := len(input)
	cardMap := make(map[int]int, maxCard-1)

	for i := 0; i < maxCard; i++ {
		cardMap[i] = 0
	}

	for cardId, rawInput := range input {
		cardMap[cardId] += 1
		rawNumbers := strings.Split(rawInput, ":")
		numbers := strings.Split(rawNumbers[1], "|")
		winningNumber := convertStringArrayToIntArray(strings.Split(numbers[0], " "))
		numbersYouHave := convertStringArrayToIntArray(strings.Split(numbers[1], " "))
		matches := getMatch(winningNumber, numbersYouHave)
		//fmt.Println(matches)
		maxCardId := cardId + matches
		if maxCardId > maxCard {
			maxCardId = maxCard
		}

		for i := cardId + 1; i <= maxCardId; i++ {
			cardMap[i] += cardMap[cardId]
			//fmt.Println(i, cardMap[i])
		}
		//fmt.Println(cardMap)
		//fmt.Println("-----")
	}

	for _, value := range cardMap {
		result += int64(value)
	}
	return result
}

func getMatch(winningNumber []int, have []int) int {
	matches := 0
	for _, number := range have {
		for _, winningNumber := range winningNumber {
			if number == winningNumber {
				matches++
			}
		}
	}
	return matches
}

func checkPrice(winningNumber []int, have []int) int64 {
	price := int64(0)
	for _, number := range have {
		for _, winningNumber := range winningNumber {
			if number == winningNumber {
				switch price {
				case 0:
					price = 1
				default:
					price *= 2
				}
			}
		}
	}
	return price
}

func convertStringArrayToIntArray(input []string) []int {
	var output []int
	for _, value := range input {
		if value == "" {
			continue
		}

		j, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		output = append(output, j)
	}
	return output
}
