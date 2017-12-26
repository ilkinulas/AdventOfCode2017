package day11

import (
	"github.com/ilkinulas/aoc2017/aoc"
	"math"
	"strings"
)

func Solution() (int, int) {
	return Solve(aoc.ReadInput("day11/input.txt"))
}

type point struct {
	x int
	y int
}

func Solve(input string) (int, int) {
	maxDistance := 0
	dirs := strings.Split(input, ",")
	center := &point{0, 0}
	p := &point{0, 0}
	for _, dir := range dirs {
		p.move(dir)
		distance := int(p.distance(center))
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return int(p.distance(&point{0, 0})), maxDistance
}

func (p *point) distance(other *point) float64 {
	//TODO is there a builtin abs function for integers?
	dx := math.Abs(float64(other.x - p.x))
	dy := math.Abs(float64(other.x - p.y))
	d := math.Abs(float64(dx - dy))

	return math.Max(math.Max(dx, dy), d)
}

func (p *point) move(dir string) {
	switch dir {
	case "n":
		p.y++
	case "s":
		p.y--
	case "ne":
		p.x++
		p.y++
	case "se":
		p.x++
	case "nw":
		p.x--
	case "sw":
		p.x--
		p.y--
	}
}
