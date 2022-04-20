package day08

import (
	"strconv"
	"strings"
)

func ApplyLine(line string, lcd [][]bool) [][]bool {
	xs := strings.Split(line, " ")
	if xs[0] == "rect" {
		x, _ := strconv.Atoi(strings.Split(xs[1], "x")[0])
		y, _ := strconv.Atoi(strings.Split(xs[1], "x")[1])

		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				lcd[i][j] = true
			}
		}
	} else if xs[0] == "rotate" {
		if xs[1] == "row" {
			y, _ := strconv.Atoi(strings.Split(xs[2], "=")[1])
			by, _ := strconv.Atoi(xs[4])

			row := make([]bool, len(lcd[y]))
			copy(row, lcd[y])

			for i := 0; i < len(row); i++ {
				lcd[y][(i+by)%len(row)] = row[i]
			}
		} else if xs[1] == "column" {
			x, _ := strconv.Atoi(strings.Split(xs[2], "=")[1])
			by, _ := strconv.Atoi(xs[4])

			col := make([]bool, len(lcd))
			for i := 0; i < len(col); i++ {
				col[i] = lcd[i][x]
			}

			for i := 0; i < len(col); i++ {
				lcd[(i+by)%len(col)][x] = col[i]
			}
		}
	}

	return lcd
}

func Solve1(lines []string) int {
	lcd := make([][]bool, 6)
	for i := 0; i < 6; i++ {
		lcd[i] = make([]bool, 50)
	}

	for _, line := range lines {
		lcd = ApplyLine(line, lcd)
	}

	count := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if lcd[i][j] {
				count++
			}
		}
	}

	return count
}

func Solve2(lines []string) string {
	lcd := make([][]bool, 6)
	for i := 0; i < 6; i++ {
		lcd[i] = make([]bool, 50)
	}

	for _, line := range lines {
		lcd = ApplyLine(line, lcd)
	}

	s := ""
	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if lcd[i][j] {
				s += "#"
			} else {
				s += " "
			}
		}
		s += "\n"
	}

	return s
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	return "Day08\nPart1: " + strconv.Itoa(Solve1(lines)) + "\nPart2:\n" + Solve2(lines)
}
