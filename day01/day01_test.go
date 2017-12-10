package day01

import "testing"

func TestSumPart1(t *testing.T) {
	tests := []struct {
		input           string
		expected        int
		nextDigitOffset int
	}{
		{"1122", 3, 1},
		{"1111", 4, 1},
		{"1234", 0, 1},
		{"91212129", 9, 1},

		{"1212", 6, 2},
		{"1221", 0, 2},
		{"123425", 4, 3},
		{"123123", 12, 3},
		{"12131415", 4, 4},
	}

	for _, test := range tests {
		sum := sum(test.input, test.nextDigitOffset)
		if sum != test.expected {
			t.Errorf("For input %s : Expected %d but found %d\n", test.input, test.expected, sum)
		}
	}

}
