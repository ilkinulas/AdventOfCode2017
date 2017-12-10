package day02

import (
	"strconv"
	"strings"

	"github.com/ilkinulas/aoc2017/aoc"
)

func Solution() (int, int) {
	input := aoc.ReadInput("day02/input.txt")
	return checksum(input, checksum1), checksum(input, checksum2)
}

func checksum1(ints []int) int {
	max, _ := aoc.Max(ints)
	min, _ := aoc.Min(ints)
	return max - min
}

func checksum2(ints []int) int {
	for i, a := range ints {
		for j, b := range ints {
			if i != j && a%b == 0 {
				return a / b
			}
		}
	}
	return 0
}

func checksum(input string, f func([]int) int) int {
	checksum := 0
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		ints := []int{}
		for _, col := range strings.Split(row, "\t") {
			i, err := strconv.Atoi(col)
			if err != nil {
				panic(err)
			}
			ints = append(ints, i)
		}
		checksum += f(ints)
	}
	return checksum
}
