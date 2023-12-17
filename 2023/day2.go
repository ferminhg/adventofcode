package adoc2023

import (
	"strconv"
	"strings"
)

func Day2(rawGames []string) (int, int) {
	total, power := 0, 0
	for id, game := range rawGames {
		splitGame := strings.Split(game, ":")
		sets := strings.Split(splitGame[1], ";")
		//fmt.Printf("Game %d: %s\n", id+1, sets)
		if isValidGame(sets) {
			total += id + 1
		}
		power += calculatePower(sets)
	}
	return total, power
}

func calculatePower(sets []string) int {
	//fmt.Printf("Sets: %s\n", sets)
	blues, reds, greens := getMaxColor(sets)
	return blues * reds * greens
}

func getMaxColor(sets []string) (int, int, int) {
	blues, reds, greens := 0, 0, 0
	for _, set := range sets {
		subset := strings.Split(set, ",")
		//fmt.Printf("Subset: %s\n", subset)
		b, r, g := countSubsetColors(subset)
		//fmt.Printf("Subset colors: %d %d %d\n", b, r, g)
		if b > blues {
			blues = b
		}
		if r > reds {
			reds = r
		}
		if g > greens {
			greens = g
		}
	}
	//fmt.Printf("Max colors: %d %d %d\n", blues, reds, greens)
	return blues, reds, greens
}

const MaxBlues = 14

const MaxReds = 12

const MaxGreens = 13

func isValid(blues int, reds int, greens int) bool {
	if blues > MaxBlues || reds > MaxReds || greens > MaxGreens {
		return false
	}
	return true
}

func isValidGame(sets []string) bool {
	for _, set := range sets {
		//fmt.Printf("Set: %s\n", set)
		subset := strings.Split(set, ",")
		b, r, g := countSubsetColors(subset)
		if !isValid(b, r, g) {
			return false
		}
	}
	return true
}

func countSubsetColors(subset []string) (int, int, int) {
	blues, reds, greens := 0, 0, 0
	for _, item := range subset {
		item = strings.Trim(item, " ")
		splitItem := strings.Split(item, " ")
		numCubes, _ := strconv.Atoi(splitItem[0])
		switch strings.Trim(splitItem[1], " ") {
		case "blue":
			blues += numCubes
		case "red":
			reds += numCubes
		case "green":
			greens += numCubes
		}
	}
	return blues, reds, greens
}
