package day06

import (
	"strconv"
	"strings"

	"github.com/ilkinulas/aoc2017/aoc"
)

func Solution() (int, int) {
	input := aoc.InputAsInts("day06/input.txt", "\t")
	return solve(input)
}

func solve(input []int) (int, int) {
	seen := make(map[string]int)
	s := intSliceToString(input)
	count := 0
	for !mapContains(seen, s) {
		seen[s] = count
		max, index := max(input)
		distribute(input, max, index)
		count++
		s = intSliceToString(input)
	}
	return count, count - seen[s]
}
func distribute(input []int, maxValue int, startIndex int) {
	size := len(input)
	input[startIndex] = 0
	for i := 0; i < maxValue; i++ {
		index := (startIndex + 1 + i) % size
		input[index] = input[index] + 1
	}
}

func mapContains(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
func intSliceToString(is []int) string {
	s := []string{}
	for i := range is {
		number := is[i]
		text := strconv.Itoa(number)
		s = append(s, text)
	}
	return strings.Join(s, ".")
}

func max(input []int) (int, int) {
	size := len(input)
	max := input[size-1]
	index := size - 1
	for i := size - 1; i >= 0; i-- {
		val := input[i]
		if val >= max {
			max = val
			index = i
		}
	}
	return max, index
}
