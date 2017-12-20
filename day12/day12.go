package day12

import (
	"github.com/ilkinulas/aoc2017/aoc"
	"strings"
)

func Solution() (int, int) {
	progMap1 := parseInput("day12/input.txt")
	progMap2 := parseInput("day12/input.txt")
	return SolvePart1(progMap1, "0"), SolvePart2(progMap2)
}

func parseInput(path string) map[string][]string {
	lines := aoc.ReadLines(path)
	progMap := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, " <-> ")
		progMap[parts[0]] = strings.Split(parts[1], ", ")
	}
	return progMap
}

func SolvePart1(programMap map[string][]string, progName string) int {
	numProgs := len(programMap)
	walkProgramMap(programMap, progName)
	return numProgs - len(programMap)
}

func SolvePart2(programMap map[string][]string) int {
	counter := 0
	for len(programMap) > 0 {
		walkProgramMap(programMap, getRandomProgram(programMap))
		counter++
	}
	return counter
}

func getRandomProgram(m map[string][]string) string {
	for k := range m {
		return k
	}
	return ""
}

func walkProgramMap(programMap map[string][]string, progName string) {
	progs, ok := programMap[progName]
	if !ok {
		return
	}
	delete(programMap, progName)
	for _, prog := range progs {
		if _, ok := programMap[prog]; !ok {
			continue
		}
		walkProgramMap(programMap, prog)
	}
}
