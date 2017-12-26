package day09

import (
	"github.com/ilkinulas/aoc2017/aoc"
	"strings"
)

func Solution() (int, int) {
	input := aoc.ReadInput("day09/input.txt")
	return calculateScore(filterInput(input)), countGarbage(input)
}

func countGarbage(input string) int {
	sum := 0
	partialSum := 0
	ignoreIndex := -1
	collect := false
	for i := 0; i < len(input); i++ {
		if i == ignoreIndex {
			continue
		}
		switch input[i] {
		case '<':
			if collect {
				partialSum++
			} else {
				collect = true
			}
		case '>':
			if collect {
				sum += partialSum
				partialSum = 0
				collect = false
			}
		case '!':
			ignoreIndex = i + 1
		default:
			if collect {
				partialSum++
			}
		}
	}
	return sum
}

func filterInput(input string) string {
	ignoreIndex := -1
	isGarbage := false
	var out []string
	for i := 0; i < len(input); i++ {
		if i == ignoreIndex {
			continue
		}
		switch input[i] {
		case '{':
			if !isGarbage {
				out = append(out, "{")
			}
		case '}':
			if !isGarbage {
				out = append(out, "}")
			}
		case '<':
			isGarbage = true
		case '>':
			isGarbage = false
		case '!':
			ignoreIndex = i + 1
		}
	}
	return strings.Join(out, "")
}

func calculateScore(s string) int {
	sum := 0
	incr := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '{' {
			sum += incr
			incr--
		} else if s[i] == '}' {
			incr++
		}
	}
	return sum
}
