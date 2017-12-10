package day03

import (
	"math"
)

const (
	UP    = iota
	LEFT  = iota
	DOWN  = iota
	RIGHT = iota
)

type grid struct {
	data   [][]int
	pX     int
	pY     int
	dir    int
	startX int
	startY int
	index  int
}

func NewGrid(maxValue int) *grid {
	size := gridSize(maxValue)
	data := make([][]int, size)

	for i := range data {
		data[i] = make([]int, size)
	}
	x := size / 2
	y := size / 2
	if size%2 == 0 {
		x--
		y--
	}
	data[x][y] = 1
	return &grid{
		data:   data,
		pX:     x,
		pY:     y,
		startX: x,
		startY: y,
		dir:    UP,
		index:  1}
}

func (g *grid) setData(value int) {
	g.data[g.pX][g.pY] = value
}

func (g *grid) getData(x int, y int) int {
	if x < 0 || y < 0 || x >= len(g.data) || y >= len(g.data) {
		return 0
	}
	return g.data[x][y]
}

func (g *grid) sumAdjacentCells(x int, y int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				sum += g.getData(x+i, y+j)
			}
		}
	}
	return sum
}

func (g *grid) next() bool {
	maxValue := len(g.data) * len(g.data)
	if g.index >= maxValue {
		return false
	}

	if g.pX == g.startX && g.pY == g.startY {
		g.pX = g.pX + 1
		g.index = g.index + 1
		return true
	}

	currentSize := gridSize(g.index)
	if g.index == currentSize*currentSize {
		g.pX, g.pY = nextXY(g.pX, g.pY, g.dir)
		g.index = g.index + 1
		currentSize = gridSize(g.index)
		g.dir = (g.dir + 1) % 4
		return true
	}

	if g.index == currentSize*currentSize-1 {
		if g.index == 3 {
			g.dir = (g.dir + 1) % 4
		}
		g.pX, g.pY = nextXY(g.pX, g.pY, g.dir)
		g.index = g.index + 1
		return true
	}

	fill := currentSize - 1
	if g.index == currentSize*currentSize-fill {
		g.dir = (g.dir + 1) % 4
	}
	g.pX, g.pY = nextXY(g.pX, g.pY, g.dir)
	g.index = g.index + 1

	return true
}

func Solution() (int, int) {
	return part1(), part2()
}

func part1() int {
	value := 265149
	g := NewGrid(value)
	for g.next() {
		if g.index == value {
			return g.pX + g.pY
		}
	}
	return 0
}
func part2() int {
	value := 265149
	g := NewGrid(value)
	g.setData(1)
	for g.next() {
		g.setData(g.sumAdjacentCells(g.pX, g.pY))
		result := g.getData(g.pX, g.pY)
		if result >= 265149 {
			return result
		}
	}
	return 0
}

func distance(grid [][]int, value int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == value {
				return i + j
			}
		}
	}
	return 0
}

func nextXY(x int, y int, direction int) (int, int) {
	switch direction {
	case RIGHT:
		x++
	case UP:
		y++
	case LEFT:
		x--
	case DOWN:
		y--
	}
	return x, y
}

func gridSize(value int) int {
	return int(math.Ceil(math.Sqrt(float64(value))))
}
