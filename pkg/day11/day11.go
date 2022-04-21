package day11

import (
	"strconv"
	"strings"

	"github.com/alexjercan/aoc-2016/pkg/util"
)

var (
	lastLevel = 4
)

type World struct {
	level      int            // The level of the elevator in the world
	generators map[string]int // The generators mapped from name to level
	microchips map[string]int // The microchips mapped from name to level
}

func (w World) String() string {
	levels := make([]string, lastLevel)

	for level := 0; level < lastLevel; level++ {
		levels[level] = "F" + strconv.Itoa(level) + " "

		if level == w.level {
			levels[level] += "E "
		} else {
			levels[level] += ". "
		}

		for gen, l := range w.generators {
			if l == level {
				levels[level] += gen + "G "
			} else {
				levels[level] += ".  "
			}
		}

		for mic, l := range w.microchips {
			if l == level {
				levels[level] += mic + "M "
			} else {
				levels[level] += ".  "
			}
		}
	}

	return strings.Join(util.Reverse(levels), "\n")
}

func parseWorld(input string) World {
	generators := make(map[string]int)
	microchips := make(map[string]int)

	for level, line := range strings.Split(strings.TrimSpace(input), "\n") {
		words := strings.Split(line, " ")

		for j := 4; j < len(words); j++ {
			if strings.HasPrefix(words[j], "generator") {
				generators[strings.ToUpper(words[j-1][:1])] = level
			} else if strings.HasPrefix(words[j], "microchip") {
				microchips[strings.ToUpper(strings.Split(words[j-1], "-")[0][:1])] = level
			}
		}
	}

	return World{
		level:      0,
		generators: generators,
		microchips: microchips,
	}
}

func moveNObjectsOneUp(n int) int {
	return (n-1)*2 - 1
}

func solve1(w World) int {
	levels := make([]int, lastLevel)

	for level := 0; level < lastLevel; level++ {
		levels[level] = 0

		for _, l := range w.generators {
			if l == level {
				levels[level]++
			}
		}

		for _, l := range w.microchips {
			if l == level {
				levels[level]++
			}
		}
	}

	currentObjects := 0
	steps := 0

	for level := 0; level < lastLevel-1; level++ {
		currentObjects += levels[level]

		steps += moveNObjectsOneUp(currentObjects)
	}

	return steps
}

func solve2(w World) int {
	w.generators["E"] = 0
	w.microchips["E"] = 0
	w.generators["D"] = 0
	w.microchips["D"] = 0

	return solve1(w)
}

func Solve(input string) string {
	world := parseWorld(input)

	return "Day11\nPart1: " + strconv.Itoa(solve1(world)) + "\nPart2: " + strconv.Itoa(solve2(world))
}
