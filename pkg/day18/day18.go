package day18

import (
	"strconv"
	"strings"
)

func step(row string) string {
	auxRow := "." + row + "."
	nextRow := ""

	for i := 0; i < len(auxRow)-2; i++ {
		if auxRow[i:i+3] == "^^." || auxRow[i:i+3] == ".^^" || auxRow[i:i+3] == "^.." || auxRow[i:i+3] == "..^" {
			nextRow += "^"
		} else {
			nextRow += "."
		}
	}

	return nextRow
}

func solve1(initialRow string) int {
	row := initialRow
	count := strings.Count(row, ".")

	for i := 1; i < 40; i++ {
		row = step(row)
		count += strings.Count(row, ".")
	}

	return count
}

func solve2(initialRow string) int {
	N := 400000

	row := initialRow
	counts := []int{strings.Count(row, ".")}
	visitedRows := make(map[string]int)
	visitedRows[row] = 0

	for i := 1; i < N; i++ {
		row = step(row)
		if index, ok := visitedRows[row]; ok {
			s1 := 0
			for j := 0; j < index; j++ {
				s1 += counts[j]
			}

			s2 := 0
			for j := index; j < i; j++ {
				s2 += counts[j]
			}

			times := (N - index) / (i - index)
			rest := (N - index) % (i - index)

			s3 := 0
			for j := index; j < index+rest; j++ {
				s3 += counts[j]
			}

			return s1 + s2*times + s3
		}
		counts = append(counts, strings.Count(row, "."))
		visitedRows[row] = i
	}

	s := 0
	for i := 0; i < N; i++ {
		s += counts[i]
	}

	return s
}

func Solve(input string) string {
	row := strings.TrimSpace(input)

	return "Day18\nPart1: " + strconv.Itoa(solve1(row)) + "\nPart2: " + strconv.Itoa(solve2(row))
}
