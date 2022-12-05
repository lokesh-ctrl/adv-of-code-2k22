package aoc22

import (
	"testing"
)

func TestFindDuplicateEntry(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   rune
	}{
		{"Abcde", "Afsj", 'A'}, {"Ljkh", "qwerL", 'L'},
	}

	for _, tt := range tests {
		testName := tt.s1 + "," + tt.s2
		t.Run(testName, func(t *testing.T) {
			result := FindDuplicateEntry(tt.s1, tt.s2)

			if result != tt.want {
				t.Error("got" + string(result) + " want" + string(tt.want))
			}
		})
	}
}

func TestFindRearrangementScore(t *testing.T) {
	var tests = []struct {
		r    rune
		want int
	}{
		{'A', 27},
		{'a', 1},
		{'z', 26},
		{'Z', 52},
	}

	for _, tt := range tests {
		testName := string(tt.r)
		t.Run(testName, func(t *testing.T) {
			result := FindRearrangementScore(tt.r)

			if result != tt.want {
				t.Errorf("expected %d, got %d", tt.want, result)
			}
		})
	}
}

func TestDay3P1(t *testing.T) {
	result := Day3P1("test.txt")
	if result != 157 {
		t.Errorf("wrong answer, got %d, want %d", result, 157)
	}
}

func TestDay3P1Log(t *testing.T) {
	result := Day3P1("input.txt")
	t.Logf("Got answer %d", result)
}
