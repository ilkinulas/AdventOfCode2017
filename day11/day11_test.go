package day11

import "testing"

func TestSimple(t *testing.T) {
	type expectation struct {
		input  string
		result int
	}

	tests := []expectation{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}

	for _, test := range tests {
		result, _ := Solve(test.input)
		if result != test.result {
			t.Errorf("Expected %d, found %d\n", test.result, result)
		}
	}
}
