package day10

import (
	"strconv"
	"strings"

	"github.com/alexjercan/aoc-2016/pkg/util"
)

func Solve1(valueToBot map[int]string, lows map[string]string, highs map[string]string) string {
	bots := make(map[string][]int)

	for value, bot := range valueToBot {
		bots[bot] = append(bots[bot], value)
	}

	for {
		for bot, values := range bots {
			if len(values) == 2 {
				if (values[0] == 61 && values[1] == 17) || (values[0] == 17 && values[1] == 61) {
					return strings.Split(bot, " ")[1]
				}

				low := lows[bot]
				high := highs[bot]

				if strings.Split(low, " ")[0] == "bot" {
					bots[low] = append(bots[low], util.Min(values[0], values[1]))
				}

				if strings.Split(high, " ")[0] == "bot" {
					bots[high] = append(bots[high], util.Max(values[0], values[1]))
				}

				delete(bots, bot)
			}
		}

		if len(bots) == 0 {
			break
		}
	}

	panic("No solution found")
}

func Solve2(valueToBot map[int]string, lows map[string]string, highs map[string]string) string {
	bots := make(map[string][]int)
	outputs := make(map[string][]int)

	for value, bot := range valueToBot {
		bots[bot] = append(bots[bot], value)
	}

	for {
		for bot, values := range bots {
			if len(values) == 2 {
				low := lows[bot]
				high := highs[bot]

				if strings.Split(low, " ")[0] == "bot" {
					bots[low] = append(bots[low], util.Min(values[0], values[1]))
				} else if strings.Split(low, " ")[0] == "output" {
					outputs[low] = append(outputs[low], util.Min(values[0], values[1]))
				}

				if strings.Split(high, " ")[0] == "bot" {
					bots[high] = append(bots[high], util.Max(values[0], values[1]))
				} else if strings.Split(high, " ")[0] == "output" {
					outputs[high] = append(outputs[high], util.Max(values[0], values[1]))
				}

				delete(bots, bot)
			}
		}

		if len(bots) == 0 {
			break
		}
	}

	return strconv.Itoa(outputs["output 0"][0] * outputs["output 1"][0] * outputs["output 2"][0])
}

func Solve(input string) string {
	valueToBot := make(map[int]string)
	lows := make(map[string]string)
	highs := make(map[string]string)

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		words := strings.Split(line, " ")

		if words[0] == "value" {
			value, _ := strconv.Atoi(words[1])
			valueToBot[value] = strings.Join(words[4:6], " ")
		} else if words[0] == "bot" {
			lows[strings.Join(words[0:2], " ")] = strings.Join(words[5:7], " ")
			highs[strings.Join(words[0:2], " ")] = strings.Join(words[10:12], " ")
		}
	}

	return "Day10\nPart1: " + Solve1(valueToBot, lows, highs) + "\nPart2: " + Solve2(valueToBot, lows, highs)
}
