package day08

import (
	"regexp"
	"github.com/ilkinulas/aoc2017/aoc"
	"strconv"
)

func Solution() (int, int) {
	re, _ := regexp.Compile("([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) (([<>=!]){1,2}) (-?[0-9]+)")
	lines := aoc.ReadLines("day08/input.txt")
	var expressions []*expr
	for _, line := range lines {
		expressions = append(expressions, parseExpression(line, re))
	}

	return solve(expressions)
}

type condition struct {
	variable string
	operator string
	value    int
}

type expr struct {
	variable  string
	amount    int
	condition *condition
}

func solve(exprs []*expr) (int, int) {
	regs := make(map[string]int)
	tmpMax:=0
	for _, expr := range exprs {
		if expr.condition.eval(regs) {
			value:=regs[expr.variable] + expr.amount
			regs[expr.variable] = value
			if value>tmpMax {
				tmpMax=value
			}
		}
	}
	max:=0
	for _, v:=range regs {
		if v>max {
			max = v
		}
	}
	return max, tmpMax
}

func parseExpression(s string, re *regexp.Regexp) *expr {
	parsed := re.FindStringSubmatch(s)
	value, _ := strconv.Atoi(parsed[7])
	condition := &condition{
		variable: parsed[4],
		operator: parsed[5],
		value:    value,
	}
	amount, _ := strconv.Atoi(parsed[3])
	if parsed[2] == "dec" {
		amount *= -1
	}
	return &expr{
		variable:  parsed[1],
		amount:    amount,
		condition: condition,
	}
}

func (c *condition) eval(register map[string]int) bool {
	value := register[c.variable]
	switch c.operator {
	case "==":
		return value == c.value
	case "!=":
		return value != c.value
	case ">":
		return value > c.value
	case "<":
		return value < c.value
	case ">=":
		return value >= c.value
	case "<=":
		return value <= c.value
	}
	return false
}
