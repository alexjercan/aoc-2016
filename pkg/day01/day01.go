package day01

import (
	"strconv"
	"strings"
)

type Command struct {
	direction string
	distance  int
}

type Position struct {
	x int
	y int
}

func ParseCommand(command string) Command {
	direction := command[0:1]
	distance, _ := strconv.Atoi(command[1:])

	return Command{direction, distance}
}

func Move(position Position, facing string, visited map[Position]int, command Command) (Position, string, map[Position]int) {
	switch facing {
	case "N":
		if command.direction == "R" {
			facing = "E"
			for i := 0; i < command.distance; i++ {
				position.x++
				visited[position]++
			}
		} else {
			facing = "W"
			for i := 0; i < command.distance; i++ {
				position.x--
				visited[position]++
			}
		}
	case "E":
		if command.direction == "R" {
			facing = "S"
			for i := 0; i < command.distance; i++ {
				position.y--
				visited[position]++
			}
		} else {
			facing = "N"
			for i := 0; i < command.distance; i++ {
				position.y++
				visited[position]++
			}
		}
	case "S":
		if command.direction == "R" {
			facing = "W"
			for i := 0; i < command.distance; i++ {
				position.x--
				visited[position]++
			}
		} else {
			facing = "E"
			for i := 0; i < command.distance; i++ {
				position.x++
				visited[position]++
			}
		}
	case "W":
		if command.direction == "R" {
			facing = "N"
			for i := 0; i < command.distance; i++ {
				position.y++
				visited[position]++
			}
		} else {
			facing = "S"
			for i := 0; i < command.distance; i++ {
				position.y--
				visited[position]++
			}
		}
	}

	return position, facing, visited
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Solve1(commands []Command) int {
	position := Position{0, 0}
	facing := "N"
	visited := make(map[Position]int)

	for _, command := range commands {
		position, facing, visited = Move(position, facing, visited, command)
	}

	return Abs(position.x) + Abs(position.y)
}

func Solve2(commands []Command) int {
	position := Position{0, 0}
	facing := "N"
	visited := make(map[Position]int)

	for _, command := range commands {
		position, facing, visited = Move(position, facing, visited, command)

		for key, value := range visited {
			if value == 2 {
				return Abs(key.x) + Abs(key.y)
			}
		}
	}

	return Abs(position.x) + Abs(position.y)
}

func Solve(input string) string {
	var commands []Command
	for _, command := range strings.Split(input, ", ") {
		commands = append(commands, ParseCommand(command))
	}

	return "Day01\nPart1: " + strconv.Itoa(Solve1(commands)) + "\nPart2: " + strconv.Itoa(Solve2(commands))
}
