package day15

import (
	"strconv"
	"strings"
)

func parseLine(line string) (int, int) {
	words := strings.Split(line, " ")

	positions, _ := strconv.Atoi(words[3])
	position, _ := strconv.Atoi(strings.Split(words[11], ".")[0])

	return position, positions
}

func solve1(mods []int, coefs []int) int {
	n := 0

	for {
		solution := true
		for i, mod := range mods {
			if (n+coefs[i])%mod != 0 {
				solution = false
				break
			}
		}

		if solution {
			return n
		}

		n++
	}
}

func solve2(mods []int, coefs []int) int {
	mods = append(mods, 11)
	coefs = append(coefs, len(coefs)+1)

	return solve1(mods, coefs)
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	mods := make([]int, 0)
	coefs := make([]int, 0)

	for i, line := range lines {
		position, positions := parseLine(line)

		mods = append(mods, positions)
		coefs = append(coefs, position+i+1)
	}

	return "Day15\nPart1: " + strconv.Itoa(solve1(mods, coefs)) + "\nPart2: " + strconv.Itoa(solve2(mods, coefs))
}
