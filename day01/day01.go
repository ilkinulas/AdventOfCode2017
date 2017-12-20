package day01

import "github.com/ilkinulas/aoc2017/aoc"

func Solution() (int, int) {
	input := aoc.ReadInput("day01/input.txt")
	return Sum(input, 1), Sum(input, len(input)/2)
}

func Sum(digits string, nextDigitOffset int) int {
	length := len(digits)
	sum := 0
	for i := 0; i < length; i++ {
		next := (i + nextDigitOffset) % length
		if digits[i] == digits[next] {
			sum += int(digits[i] - '0')
		}
	}
	return sum
}
