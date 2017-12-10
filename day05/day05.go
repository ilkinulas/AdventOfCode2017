package day05

import (
	"github.com/ilkinulas/aoc2017/aoc"
)

func Solution() (int, int) {
	input := aoc.InputAsInts("day05/input.txt", "\n")
	part1 := solve(input, func(offset int) int { return offset + 1 })

	input = aoc.InputAsInts("day05/input.txt", "\n")
	part2 := solve(input, func(offset int) int {
		if offset >= 3 {
			return offset - 1
		}
		return offset + 1
	})
	return part1, part2
}

func solve(offsets []int, offsetUpdater func(int) int) int {
	index := 0
	size := len(offsets)
	count := 0
	for index < size {
		offset := offsets[index]
		jumpTo := nextIndex(index, offset, size)
		if jumpTo >= 0 {
			offsets[index] = offsetUpdater(offset)
			count++
			index = jumpTo
		} else {
			count++
			return count
		}
	}
	return 0
}

func nextIndex(currentIndex int, offset int, size int) int {
	next := currentIndex + offset
	if next >= size || next < 0 {
		return -1
	}
	return next
}
