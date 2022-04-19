package day02

import (
	"strings"

	"github.com/alexjercan/aoc-2016/pkg/util"
)

var (
	keypad = [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	keypadB = [][]string{
		{"0", "0", "1", "0", "0"},
		{"0", "2", "3", "4", "0"},
		{"5", "6", "7", "8", "9"},
		{"0", "A", "B", "C", "0"},
		{"0", "0", "D", "0", "0"},
	}
)

func Move(direction string, x int, y int) (int, int) {
	switch direction {
	case "U":
		y--
	case "D":
		y++
	case "L":
		x--
	case "R":
		x++
	}

	x = util.Clamp(x, 0, 2)
	y = util.Clamp(y, 0, 2)

	return x, y
}

func MoveB(direction string, x int, y int) (int, int) {
	switch direction {
	case "U":
		if y > 0 && keypadB[y-1][x] != "0" {
			y--
		}
	case "D":
		if y < 4 && keypadB[y+1][x] != "0" {
			y++
		}
	case "L":
		if x > 0 && keypadB[y][x-1] != "0" {
			x--
		}
	case "R":
		if x < 4 && keypadB[y][x+1] != "0" {
			x++
		}
	}

	return x, y
}

func Solve1(moves []string) string {
	x := 1
	y := 1

	keycode := ""

	for _, move := range moves {
		for _, direction := range move {
			x, y = Move(string(direction), x, y)
		}

		keycode += keypad[y][x]
	}

	return keycode
}

func Solve2(moves []string) string {
	x := 0
	y := 2

	keycode := ""

	for _, move := range moves {
		for _, direction := range move {
			x, y = MoveB(string(direction), x, y)
		}

		keycode += keypadB[y][x]
	}

	return keycode
}

func Solve(input string) string {
	moves := strings.Split(strings.TrimSpace(input), "\n")

	return "Day02\nPart1: " + Solve1(moves) + "\nPart2: " + Solve2(moves)
}
