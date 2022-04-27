package day23

import (
	"strconv"
	"strings"
)

func step(i int, registers map[string]int, instructions [][]string) (int, map[string]int) {

	if i+5 < len(instructions) && instructions[i+0][0] == "inc" && instructions[i+1][0] == "dec" && instructions[i+2][0] == "jnz" && instructions[i+2][1] == instructions[i+1][1] && instructions[i+2][2] == "-2" && instructions[i+3][0] == "dec" && instructions[i+4][0] == "jnz" && instructions[i+4][1] == instructions[i+3][1] && instructions[i+4][2] == "-5" {
		X := instructions[i+0][1]
		Y := instructions[i+1][1]
		Z := instructions[i+3][1]

		registers[X] = registers[Y] * registers[Z]
		registers[Y] = 0
		registers[Z] = 0

		i += 5
	} else {
		instruction := instructions[i]

		switch instruction[0] {
		case "cpy":
			if v, err := strconv.Atoi(instruction[1]); err == nil {
				registers[instruction[2]] = v
			} else {
				registers[instruction[2]] = registers[instruction[1]]
			}
			i++
		case "inc":
			registers[instruction[1]]++
			i++
		case "dec":
			registers[instruction[1]]--
			i++
		case "jnz":
			value := 0
			if v, err := strconv.Atoi(instruction[1]); err == nil {
				value = v
			} else {
				value = registers[instruction[1]]
			}

			if value != 0 {
				jvalue := 0
				if jv, err := strconv.Atoi(instruction[2]); err == nil {
					jvalue = jv
				} else {
					jvalue = registers[instruction[2]]
				}

				i += jvalue
			} else {
				i++
			}
		case "tgl":
			index := 0
			if v, err := strconv.Atoi(instruction[1]); err == nil {
				index = i + v
			} else {
				index = i + registers[instruction[1]]
			}

			if index >= 0 && index < len(instructions) {
				tgl := instructions[index]
				switch tgl[0] {
				case "inc":
					tgl[0] = "dec"
				case "dec":
					tgl[0] = "inc"
				case "tgl":
					tgl[0] = "inc"
				case "jnz":
					tgl[0] = "cpy"
				case "cpy":
					tgl[0] = "jnz"
				}
			}

			i++
		}
	}

	return i, registers
}

func solve1(lines []string) int {
	instructions := make([][]string, len(lines))
	for i, line := range lines {
		instructions[i] = strings.Split(line, " ")
	}

	registers := make(map[string]int)
	registers["a"] = 7
	i := 0
	for {
		i, registers = step(i, registers, instructions)
		if i < 0 || i >= len(instructions) {
			break
		}
	}
	return registers["a"]
}

func solve2(lines []string) int {
	instructions := make([][]string, len(lines))
	for i, line := range lines {
		instructions[i] = strings.Split(line, " ")
	}

	registers := make(map[string]int)
	registers["a"] = 12
	i := 0
	for {
		i, registers = step(i, registers, instructions)
		if i < 0 || i >= len(instructions) {
			break
		}
	}
	return registers["a"]
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	return "Day23\nPart1: " + strconv.Itoa(solve1(lines)) + "\nPart2: " + strconv.Itoa(solve2(lines))
}
