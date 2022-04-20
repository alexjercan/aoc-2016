package day09

import (
	"strconv"
	"strings"
)

func Solve1(input string) int {
	count := 0
	i := 0

	for {
		if i >= len(input) {
			break
		}

		if input[i] == '(' {
			depth := 0
			times := 0

			for j := i + 1; j < len(input); j++ {
				if input[j] == 'x' {
					depth, _ = strconv.Atoi(input[i+1 : j])
					i = j
				} else if input[j] == ')' {
					times, _ = strconv.Atoi(input[i+1 : j])
					i = j + 1
					break
				}
			}

			count += depth * times
			i += depth
		} else {
			count++
			i++
		}
	}

	return count
}

func Solve2(input string) int {
	count := 0
	i := 0

	for {
		if i >= len(input) {
			break
		}

		if input[i] == '(' {
			depth := 0
			times := 0

			for j := i + 1; j < len(input); j++ {
				if input[j] == 'x' {
					depth, _ = strconv.Atoi(input[i+1 : j])
					i = j
				} else if input[j] == ')' {
					times, _ = strconv.Atoi(input[i+1 : j])
					i = j + 1
					break
				}
			}

			count += Solve2(input[i:i+depth]) * times
			i += depth
		} else {
			count++
			i++
		}
	}

	return count
}

func Solve(input string) string {
	input = strings.TrimSpace(input)

	return "Day09\nPart1: " + strconv.Itoa(Solve1(input)) + "\nPart2: " + strconv.Itoa(Solve2(input))
}
