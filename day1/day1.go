package aoc22

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func Day1p1(inputFile string) int {
	eachElfCalories := GetEachElfCalories(inputFile)
	sort.Ints(eachElfCalories)
	return eachElfCalories[len(eachElfCalories)-1]
}

func Day2p2(inputFile string) int {
	eachElfCalories := GetEachElfCalories(inputFile)
	sort.Ints(eachElfCalories)
	lengthOfArray := len(eachElfCalories)
	return eachElfCalories[lengthOfArray-1] + eachElfCalories[lengthOfArray-2] + eachElfCalories[lengthOfArray-3]
}

func parseAndReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetEachElfCalories(inputFile string) []int {
	var eachElfCalories []int
	currentElfCalories := 0

	calories, _ := parseAndReadLines(inputFile)

	for _, value := range calories {
		if value != "" {
			n, _ := strconv.Atoi(value)
			currentElfCalories += n
		} else {
			eachElfCalories = append(eachElfCalories, currentElfCalories)
			currentElfCalories = 0
		}
	}
	return eachElfCalories
}
