package main

import (
	"testing"
	"fmt"
	"github.com/ilkinulas/aoc2017/day01"
	"github.com/ilkinulas/aoc2017/day02"
	"github.com/ilkinulas/aoc2017/day03"
	"github.com/ilkinulas/aoc2017/day04"
	"github.com/ilkinulas/aoc2017/day05"
	"github.com/ilkinulas/aoc2017/day06"
	"github.com/ilkinulas/aoc2017/day07"
	"github.com/ilkinulas/aoc2017/day08"
	"github.com/ilkinulas/aoc2017/day09"
	"github.com/ilkinulas/aoc2017/day10"
	"github.com/ilkinulas/aoc2017/day11"
	"github.com/ilkinulas/aoc2017/day12"
)

func TestDay01(t *testing.T) {
	part1, part2 := day01.Solution()
	assertEqual(t, 995, part1)
	assertEqual(t, 1130, part2)
}

func TestDay02(t *testing.T) {
	part1, part2 := day02.Solution()
	assertEqual(t, 53460, part1)
	assertEqual(t, 282, part2)
}

func TestDay03(t *testing.T) {
	part1, part2 := day03.Solution()
	assertEqual(t, 438, part1)
	assertEqual(t, 266330, part2)
}

func TestDay04(t *testing.T) {
	part1, part2 := day04.Solution()
	assertEqual(t, 386, part1)
	assertEqual(t, 208, part2)
}

func TestDay05(t *testing.T) {
	part1, part2 := day05.Solution()
	assertEqual(t, 359348, part1)
	assertEqual(t, 27688760, part2)
}

func TestDay06(t *testing.T) {
	part1, part2 := day06.Solution()
	assertEqual(t, 5042, part1)
	assertEqual(t, 1086, part2)
}

func TestDay07(t *testing.T) {
	part1, part2 := day07.Solution()
	assertEqual(t, "bpvhwhh", part1)
	assertEqual(t, 256, part2)
}

func TestDay08(t *testing.T) {
	part1, part2 := day08.Solution()
	assertEqual(t, 4448, part1)
	assertEqual(t, 6582, part2)
}

func TestDay09(t *testing.T) {
	part1, part2 := day09.Solution()
	assertEqual(t, 13154, part1)
	assertEqual(t, 6369, part2)
}

func TestDay10(t *testing.T) {
	part1, part2 := day10.Solution()
	assertEqual(t, 4114, part1)
	assertEqual(t, "2f8c3d2100fdd57cec130d928b0fd2dd", part2)
}
func TestDay11(t *testing.T) {
	part1, part2 := day11.Solution()
	assertEqual(t, 743, part1)
	assertEqual(t, 1493, part2)
}

func TestDay12(t *testing.T) {
	part1, part2 := day12.Solution()
	assertEqual(t, 175, part1)
	assertEqual(t, 213, part2)
}

func assertEqual(t *testing.T, expected interface{}, value interface{}) {
	if expected == value {
		return
	}
	message := fmt.Sprintf("%v != %v", expected, value)
	t.Fatal(message)
}
