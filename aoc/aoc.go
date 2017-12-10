package aoc

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(b)
}

func ReadLines(file string) []string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return strings.Split(string(b), "\n")
}

func InputAsInts(file string, sep string) []int {
	input := ReadInput(file)
	parts := strings.Split(input, sep)
	ints := []int{}
	for _, part := range parts {
		intVal, _ := strconv.Atoi(part)
		ints = append(ints, intVal)
	}
	return ints
}

func Max(ints []int) (int, int) {
	max := ints[0]
	index := 0
	for i, val := range ints {
		if val > max {
			max = val
			index = i
		}
	}
	return max, index
}

func Min(ints []int) (int, int) {
	min := ints[0]
	index := 0
	for i, val := range ints {
		if val < min {
			min = val
			index = i
		}
	}
	return min, index
}
