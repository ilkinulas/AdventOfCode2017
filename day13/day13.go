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
	for d, r := range f.data {
		if caught(r, d) {
			severity += d * r
		}
	}
	return severity
}

func (f *Firewall) packetDelay() (delay int) {
scan:
	for {
		for d, r := range f.data {
			if caught(r, d+delay) {
				delay++
				continue scan
			}
		}
		return delay
	}
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
