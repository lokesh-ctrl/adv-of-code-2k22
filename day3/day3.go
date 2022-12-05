package aoc22

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func Day3P1(inputFile string) int {
	inputs, _ := parseAndReadLines(inputFile)
	total := 0

	for _, input := range inputs {
		inputArray := []string{input[0 : len(input)/2], input[len(input)/2:]}
		total += FindRearrangementScore(FindDuplicateEntry(inputArray)[0])
	}

	return total
}

func Day3P2(inputFile string) int {
	inputs, _ := parseAndReadLines(inputFile)
	total := 0

	for i := 0; i < len(inputs); i += 3 {
		total += FindRearrangementScore(FindDuplicateEntry(inputs[i : i+3])[0])
	}

	return total
}

func FindRearrangementScore(duplicate rune) int {
	scores := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
	}
	if value, ok := scores[duplicate]; ok {
		return value
	} else {
		return scores[unicode.ToLower(duplicate)] + 26
	}
}

func FindDuplicateEntry(stringArray []string) []rune {
	seenChars := []rune(stringArray[0])
	for i := 1; i < len(stringArray); i++ {
		var newChars = []rune{}
		for _, char := range seenChars {
			if strings.Contains(stringArray[i], string(char)) {
				newChars = append(newChars, char)
			}
		}
		seenChars = newChars
	}
	return seenChars
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
