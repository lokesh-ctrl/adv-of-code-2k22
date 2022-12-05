package aoc22

import (
	"testing"
)

func TestFindDuplicateEntry(t *testing.T) {
	var tests = []struct {
		strings []string
		want    rune
	}{
		{
			strings: []string{"Abcde", "aaaaaA"},
			want:    'A',
		},
		{
			strings: []string{"Lasdf", "Lasdf", "Llll"},
			want:    'L',
		},
		{
			strings: []string{"Lasdf", "Loki", "L"},
			want:    'L',
		},
	}
	for _, tt := range tests {
		testName := tt.want

		t.Run(string(testName), func(t *testing.T) {
			result := FindDuplicateEntry(tt.strings)

			if result[0] != tt.want {
				t.Error("got" + string(result) + "want" + string(tt.want))
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

func TestDay3P2(t *testing.T) {
	result := Day3P2("test.txt")
	if result != 70 {
		t.Errorf("wrong answer, got %d, want %d", result, 70)
	}
}

func TestDay3P2Log(t *testing.T) {
	result := Day3P2("input.txt")
	t.Logf("Got answer %d", result)
}
