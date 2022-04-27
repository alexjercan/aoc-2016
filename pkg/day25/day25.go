package day25

import (
	"strconv"
	"strings"
)

func step(a int, b int, c int) (int, bool) {
	seq := make([]int, 0)
	og := a

	d := a + c*b
	for i := 0; i < 1000; i++ {
		a = d

		for {
			c = 2 - (a % 2)
			a = a / 2

			b = 2 - c

			seq = append(seq, b)

			if a == 0 {
				break
			}
		}
	}

	ok := true
	for i := 0; i < len(seq); i += 2 {
		if seq[i] != 0 || seq[i+1] != 1 {
			ok = false
		}
	}

	return og, ok
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := make([][]string, len(lines))
	for i, line := range lines {
		instructions[i] = strings.Split(line, " ")
	}

	for i := 0; i < 700; i++ {
		b, _ := strconv.Atoi(instructions[2][1])
		c, _ := strconv.Atoi(instructions[1][1])
		og, ok := step(i, b, c)

		if ok {
			return "Day23\nPart1: " + strconv.Itoa(og)
		}
	}

	return "Day23\nPart1: nan"
}
