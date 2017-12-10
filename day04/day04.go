package day04

import (
	"sort"
	"strings"

	"github.com/ilkinulas/aoc2017/aoc"
)

func Solution() (int, int) {
	input := aoc.ReadInput("day04/input.txt")
	return numValid(input)
}

func numValid(input string) (int, int) {
	lines := strings.Split(input, "\n")
	countNoDups := 0
	countNoAnagrams := 0
	for _, line := range lines {
		words := strings.Fields(line)
		if isValid(words, func(s string) string { return s }) {
			countNoDups++
		}
		if isValid(words, func(s string) string { return sortString(s) }) {
			countNoAnagrams++
		}
	}
	return countNoDups, countNoAnagrams
}

func isValid(words []string, f func(string) string) bool {
	m := make(map[string]int)
	for _, word := range words {
		word = f(word)
		if _, ok := m[word]; ok {
			return false
		}
		m[word] = 1
	}
	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
