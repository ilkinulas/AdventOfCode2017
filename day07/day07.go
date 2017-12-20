package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ilkinulas/aoc2017/aoc"
)

type program struct {
	name        string
	weight      int
	totalWeight int
	children    []*program
	parent      *program
}

func Solution() (string, int) {
	lines := aoc.ReadLines("day07/input.txt")
	re := regexp.MustCompile("([a-z]+) \\(([0-9]+)\\)( -> (.*)){0,1}")
	programMap := make(map[string]*program)
	for _, line := range lines {
		program := newProgram(line, re)
		programMap[program.name] = program
	}
	root := findRoot(programMap)
	prepareTree(root, programMap)

	return root.name, findBalancedWeight(root)
}

func prepareTree(root *program, input map[string]*program) int {
	if len(root.children) == 0 {
		return root.weight
	}
	for i := range root.children {
		child, _ := input[root.children[i].name]
		root.children[i] = child
		root.children[i].parent = root
		root.totalWeight += prepareTree(child, input)
	}
	return root.totalWeight
}

func traverseTree(root *program) {
	println(root.name, root.weight, root.totalWeight)
	if len(root.children) == 0 {
		return
	}
	for i := range root.children {
		traverseTree(root.children[i])

	}
}

func findBalancedWeight(root *program) int {
	prog := findUnbalanced(root)
	if prog != root && prog != nil {
		ub := findUnbalanced(prog)
		if ub == prog || ub == nil {
			diff := 0
			for _, sibling := range prog.parent.children {
				diff = prog.totalWeight - sibling.totalWeight
				if diff != 0 {
					break
				}
			}
			return prog.weight - diff
		}
		return findBalancedWeight(ub)
	}
	return 0
}

func findUnbalanced(root *program) *program {
	if len(root.children) == 0 {
		return root
	}

	seen := make(map[int]int)
	for i := range root.children {
		weight := root.children[i].totalWeight
		if val, ok := seen[weight]; ok {
			seen[weight] = val + 1
		} else {
			seen[weight] = 1
		}
	}
	for k, v := range seen {
		if v == 1 {
			for i := range root.children {
				if root.children[i].totalWeight == k {
					return root.children[i]
				}
			}
		}
	}
	return nil
}

func findRoot(programMap map[string]*program) *program {
	childMap := make(map[string]bool)
	for _, v := range programMap {
		for i := range v.children {
			childMap[v.children[i].name] = true
		}
	}
	for name, prog := range programMap {
		if _, ok := childMap[name]; !ok {
			return prog
		}
	}
	return nil
}

func (p *program) print() {
	fmt.Printf("%s - %d - %v\n", p.name, p.weight, p.children)
}

func newProgram(line string, re *regexp.Regexp) *program {
	result := re.FindStringSubmatch(line)
	weight, _ := strconv.Atoi(result[2])
	childNames := strings.Split(result[4], ", ")
	if childNames[0] == "" {
		childNames = []string{}
	}
	children := []*program{}
	for i := range childNames {
		children = append(children, &program{name: childNames[i]})
	}
	p := &program{
		name:        result[1],
		weight:      weight,
		children:    children,
		totalWeight: weight,
	}
	return p
}
