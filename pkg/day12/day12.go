package day12

import (
	"strconv"
	"strings"
)

type Registers map[string]int

func step(i int, registers Registers, instructions [][]string) (int, Registers) {
	switch instructions[i][0] {
	case "cpy":
		if v, err := strconv.Atoi(instructions[i][1]); err == nil {
			registers[instructions[i][2]] = v
		} else {
			registers[instructions[i][2]] = registers[instructions[i][1]]
		}
		i++
	case "inc":
		registers[instructions[i][1]]++
		i++
	case "dec":
		registers[instructions[i][1]]--
		i++
	case "jnz":
		if jv, err := strconv.Atoi(instructions[i][2]); err == nil {
			if v, err := strconv.Atoi(instructions[i][1]); err == nil {
				if v != 0 {
					i += jv
				} else {
					i++
				}
			} else {
				if registers[instructions[i][1]] != 0 {
					i += jv
				} else {
					i++
				}
			}
		}
	}

	return i, registers
}

func solve1(instructions [][]string) int {
	registers := make(Registers)
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 0
	registers["d"] = 0

	i := 0
	for i < len(instructions) {
		i, registers = step(i, registers, instructions)
	}

	return registers["a"]
}

func solve2(instructions [][]string) int {
	registers := make(Registers)
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 1
	registers["d"] = 0

	i := 0
	for i < len(instructions) {
		i, registers = step(i, registers, instructions)
	}

	return registers["a"]
}

func Solve(input string) string {
	instructions := make([][]string, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		instructions = append(instructions, strings.Split(line, " "))
	}

	return "Day12\nPart1: " + strconv.Itoa(solve1(instructions)) + "\nPart2: " + strconv.Itoa(solve2(instructions))
}
