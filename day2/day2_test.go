package aoc22

import (
	"fmt"
	"testing"
)

func TestFindTheScoreWonForGameForWinningMatch(t *testing.T) {
	result := findTheScoreWonForGame("A Y")
	if result != 8 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestFindTheScoreWonForGameForLossingMatch(t *testing.T) {
	result := findTheScoreWonForGame("B X")
	if result != 1 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestFindTheScoreWonForGameForDrawMatch(t *testing.T) {
	result := findTheScoreWonForGame("C Z")
	if result != 6 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestDay2Part1Test(t *testing.T) {
	result := Day2P1("test.txt")

	if result != 15 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestDay2Part1ResultLog(t *testing.T) {
	result := Day2P1("input.txt")
	fmt.Printf("t: %v\n", result)
}

func TestFindTheScoreForResultForDrawMatch(t *testing.T) {
	result := findTheScoreForResult("A Y")
	if result != 4 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestFindTheScoreForResultForLoseMatch(t *testing.T) {
	result := findTheScoreForResult("B X")
	if result != 1 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestFindTheScoreForResultForWinMatch(t *testing.T) {
	result := findTheScoreForResult("C Z")
	if result != 7 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestDay2Part2Test(t *testing.T) {
	result := Day2P2("test.txt")

	if result != 12 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestDay2Part2ResultLog(t *testing.T) {
	result := Day2P2("input.txt")
	fmt.Printf("t: %v\n", result)
}
