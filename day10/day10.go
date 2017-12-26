package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution() (int, string) {
	puzzleInput := "165,1,255,31,87,52,24,113,0,91,148,254,158,2,73,153"
	return solvePart1(puzzleInput), solvePart2(puzzleInput)
}

func solvePart1(s string) int {
	puzzleInput := strings.Split(s, ",")
	var input []int
	for i := 0; i < len(puzzleInput); i++ {
		val, _ := strconv.Atoi(puzzleInput[i])
		input = append(input, val)
	}
	var ints []int
	for i := 0; i < 256; i++ {
		ints = append(ints, i)
	}
	skip := 0
	pos := 0
	for i := range input {
		reverse(ints, pos, input[i])
		pos = (input[i] + pos + skip) % len(ints)
		skip = skip + 1
	}
	return ints[0] * ints[1]
}

func solvePart2(s string) string {
	var lengths []int
	for i := range s {
		lengths = append(lengths, int(rune(s[i])))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	var ints []int
	for i := 0; i < 256; i++ {
		ints = append(ints, i)
	}
	skip := 0
	pos := 0
	for round := 0; round < 64; round++ {
		for i := range lengths {
			reverse(ints, pos, lengths[i])
			pos = (lengths[i] + pos + skip) % len(ints)
			skip = skip + 1
		}
	}
	var xors []int
	for i := 0; i < 16; i++ {
		xors = append(xors, xor(ints[i*16:i*16+16]))
	}

	var hash []string
	for i := range xors {
		hash = append(hash, fmt.Sprintf("%02x", xors[i]))
	}
	return strings.Join(hash, "")
}

func xor(ints []int) int {
	result := 0
	for i := range ints {
		result = result ^ ints[i]
	}
	return result
}

//TODO find a more efficient way of reversing a part of a slice
func reverse(ints []int, from int, length int) {
	var reversed []int
	for i := length - 1; i >= 0; i-- {
		index := (i + from) % len(ints)
		reversed = append(reversed, ints[index])
	}

	for i := 0; i < length; i++ {
		index := (i + from) % len(ints)
		ints[index] = reversed[i]
	}
}
