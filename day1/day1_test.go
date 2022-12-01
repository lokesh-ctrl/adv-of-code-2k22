package aoc22

import (
	"fmt"
	"testing"
)

func TestDay1Part1Test(t *testing.T) {
	result := Day1p1("test.txt")

	if result != 24000 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestDay1Part1ResultLog(t *testing.T) {
	result := Day1p1("input.txt")
	fmt.Printf("t: %v\n", result)
}

func TestDay1Part2Test(t *testing.T) {
	result := Day1p1("input.txt")
	if result != 45000 {
		t.Errorf("wrong answer got %d", result)
	}
}

func TestDay1Part2ResultLog(t *testing.T) {
	result := Day2p2("input.txt")
	fmt.Printf("t: %v\n", result)
}
