package adv22

import "testing"

func TestFindIfOverlapping(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{{"2-4,6-8", false}, {"2-3,4-5", false}, {"5-7,7-9", false}, {"2-8,3-7", true}, {"6-6,4-6", true}, {"2-6,4-8", false}, {"28-96,28-28", true}}

	for _, tt := range tests {
		testName := tt.input
		t.Run(testName, func(t *testing.T) {
			result := FindIfOverallOverlapping(tt.input)
			if result != tt.output {
				t.Error("wrong answer")
			}
		})
	}
}

func TestDay4P1(t *testing.T) {
	result := Day4P1("test.txt")
	if result != 2 {
		t.Errorf("wrong answer, got %d, want %d", result, 2)
	}
}

func TestDay4P1Log(t *testing.T) {
	result := Day4P1("input.txt")
	t.Logf("Got answer %d", result)
}
