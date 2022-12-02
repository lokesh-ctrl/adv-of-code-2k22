package aoc22

import (
	"bufio"
	"os"
)

func Day2P1(inputFile string) int {
	eachGameScores, _ := parseAndReadLines(inputFile)
	totalScore := 0
	for _, value := range eachGameScores {
		totalScore += findTheScoreWonForGame(value)
	}

	return totalScore
}

func Day2P2(inputFile string) int {
	eachGameScores, _ := parseAndReadLines(inputFile)
	totalScore := 0
	for _, value := range eachGameScores {
		totalScore += findTheScoreForResult(value)
	}

	return totalScore
}

func findTheScoreForResult(gameInput string) int {
	score := 0
	winRules := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	drawRules := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	loseRules := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	optionScores := map[string]int{"X": 1, "Y": 2, "Z": 3}
	if string(gameInput[2]) == "Z" {
		score += 6
		score += optionScores[winRules[string(gameInput[0])]]
	} else if string(gameInput[2]) == "Y" {
		score += 3
		score += optionScores[drawRules[string(gameInput[0])]]
	} else {
		score += optionScores[loseRules[string(gameInput[0])]]
	}
	return score
}

func findTheScoreWonForGame(gameInput string) int {
	score := 0
	winRules := map[string]string{"X": "C", "Y": "A", "Z": "B"}
	drawRules := map[string]string{"X": "A", "Y": "B", "Z": "C"}
	optionScores := map[string]int{"X": 1, "Y": 2, "Z": 3}
	if winRules[string(gameInput[2])] == string(gameInput[0]) {
		score += 6
	} else if string(gameInput[0]) == drawRules[string(gameInput[2])] {
		score += 3
	}
	score += optionScores[string(gameInput[2])]
	return score
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
