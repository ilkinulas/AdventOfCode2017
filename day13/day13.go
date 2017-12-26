package day13

import (
	"github.com/ilkinulas/aoc2017/aoc"
	"strconv"
	"strings"
)

func Solution() (int, int) {
	firewall := newFirewall("day13/input.txt")
	return firewall.severity(), firewall.packetDelay()
}

type Firewall struct {
	data map[int]int
}

func (f *Firewall) severity() (severity int) {
	for depth, iRange := range f.data {
		if caught(iRange, depth) {
			severity += depth * iRange
		}
	}
	return severity
}

func (f *Firewall) packetDelay() (delay int) {
	for f.caughtWhilePassingThrough(delay) {
		delay++
	}
	return delay
}

func (f*Firewall) caughtWhilePassingThrough(delay int) bool {
	for depth, iRange := range f.data {
		if caught(iRange, depth+delay) {
			return true
		}
	}
	return false
}

func caught(sRange int, delay int) bool {
	if sRange == 1 {
		return true
	}
	return delay%(2*(sRange-1)) == 0
}

func newFirewall(input string) Firewall {
	firewall := Firewall{data: make(map[int]int)}
	for _, line := range aoc.ReadLines(input) {
		parts := strings.Split(line, ": ")
		depth, _ := strconv.Atoi(parts[0])
		sRange, _ := strconv.Atoi(parts[1])
		firewall.data[depth] = sRange
	}
	return firewall
}
