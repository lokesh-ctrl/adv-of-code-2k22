package adv22

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day4P1(inputFile string) int {
	inputs, _ := parseAndReadLines(inputFile)
	numOLap := 0
	for _, input := range inputs {
		if FindIfOverallOverlapping(input) {
			numOLap += 1
		}
	}
	return numOLap
}

func FindIfOverallOverlapping(s string) bool {
	ranges := strings.Split(s, ",")
	startingOfFirst, _ := strconv.Atoi(strings.Split(ranges[0], "-")[0])
	endingOfFirst, _ := strconv.Atoi(strings.Split(ranges[0], "-")[1])
	startingOfSecond, _ := strconv.Atoi(strings.Split(ranges[1], "-")[0])
	endingOfSecond, _ := strconv.Atoi(strings.Split(ranges[1], "-")[1])
	if (startingOfFirst <= startingOfSecond && endingOfSecond <= endingOfFirst) || (startingOfFirst >= startingOfSecond && endingOfSecond >= endingOfFirst) {
		return true
	} else {
		return false
	}
}

func FindIfAnyOverlap(s string) bool {
	ranges := strings.Split(s, ",")
	startingOfFirst, _ := strconv.Atoi(strings.Split(ranges[0], "-")[0])
	endingOfFirst, _ := strconv.Atoi(strings.Split(ranges[0], "-")[1])
	startingOfSecond, _ := strconv.Atoi(strings.Split(ranges[1], "-")[0])
	endingOfSecond, _ := strconv.Atoi(strings.Split(ranges[1], "-")[1])
	if math.Max(float64(startingOfFirst), float64(startingOfSecond)) <= math.Min(float64(endingOfFirst), float64(endingOfSecond)) {
		return true
	} else {
		return false
	}
}

func Day4P2(inputFile string) int {
	inputs, _ := parseAndReadLines(inputFile)
	numOLap := 0
	for _, input := range inputs {
		if FindIfAnyOverlap(input) {
			numOLap += 1
		}
	}
	return numOLap
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
