package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ilkinulas/aoc2017/aoc"
)

type program struct {
	name     string
	weight   int
	children []string
}

func (p *program) print() {
	fmt.Printf("Program %s - %d - %v\n", p.name, p.weight, p.children)
}

func newProgram(line string, re *regexp.Regexp) *program {
	result := re.FindStringSubmatch(line)
	weight, _ := strconv.Atoi(result[2])
	children := strings.Split(result[4], ", ")
	if children[0] == "" {
		children = []string{}
	}
	p := &program{
		name:     result[1],
		weight:   weight,
		children: children,
	}
	return p
}

func Solution() (string, int) {
	lines := aoc.ReadLines("day07/input.txt")
	re, _ := regexp.Compile("([a-z]+) \\(([0-9]+)\\)( -> (.*)){0,1}")
	programMap := make(map[string]*program)
	for _, line := range lines {
		program := newProgram(line, re)
		programMap[program.name] = program
	}
	return findRoot(programMap).name, 0
}

func findRoot(programMap map[string]*program) *program {
	childMap := make(map[string]bool)
	for _, v := range programMap {
		for i := range v.children {
			childMap[v.children[i]] = true
		}
	}
	for name, prog := range programMap {
		if _, ok := childMap[name]; !ok {
			return prog
		}
	}
	return nil
}
