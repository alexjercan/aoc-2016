package day06

import (
	"strings"

	"github.com/alexjercan/aoc-2016/pkg/util"
)

func Solve1(codes [][]rune) []rune {
	result := make([]rune, len(codes))

	for _, code := range codes {
		runeCount := make(map[rune]int)
		for _, r := range code {
			runeCount[r]++
		}
		var max rune
		var maxCount int
		for r, count := range runeCount {
			if count > maxCount {
				max = r
				maxCount = count
			}
		}
		result = append(result, max)
	}

	return result
}

func Solve2(codes [][]rune) []rune {
	result := make([]rune, len(codes))

	for _, code := range codes {
		runeCount := make(map[rune]int)
		for _, r := range code {
			runeCount[r]++
		}
		var min rune
		var minCount int
		for r, count := range runeCount {
			if count < minCount || minCount == 0 {
				min = r
				minCount = count
			}
		}
		result = append(result, min)
	}

	return result
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	chars := make([][]rune, len(lines))
	for i, line := range lines {
		chars[i] = []rune(line)
	}
	codes := util.Transpose(chars)

	return "Day06\nPart1: " + string(Solve1(codes)) + "\nPart2: " + string(Solve2(codes))
}
