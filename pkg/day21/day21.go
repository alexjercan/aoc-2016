package day21

import (
	"strconv"
	"strings"
)

func swapPosition(str string, i int, j int) string {
	if i > j {
		i, j = j, i
	}

	a := str[i]
	b := str[j]
	return str[:i] + string(b) + str[i+1:j] + string(a) + str[j+1:]
}

func swapLetter(str string, a string, b string) string {
	newStr := ""

	for _, c := range str {
		if string(c) == a {
			newStr += b
		} else if string(c) == b {
			newStr += a
		} else {
			newStr += string(c)
		}
	}

	return newStr
}

func rotateBased(password string, letter string) string {
	index := strings.Index(password, letter)
	if index >= 4 {
		index++
	}
	index++
	index %= len(password)
	return rotate(password, "right", index)
}

func rotate(str string, direction string, steps int) string {
	steps %= len(str)

	if direction == "left" {
		return str[steps:] + str[:steps]
	} else {
		return str[len(str)-steps:] + str[:len(str)-steps]
	}
}

func reverse(str string, a int, b int) string {
	xs := strings.Split(str, "")
	newStr := ""

	for i := 0; i < a; i++ {
		newStr += xs[i]
	}

	for i := b; i >= a; i-- {
		newStr += xs[i]
	}

	for i := b + 1; i < len(xs); i++ {
		newStr += xs[i]
	}

	return newStr
}

func move(str string, i int, j int) string {
	a := str[i]
	str = str[:i] + str[i+1:]
	return str[:j] + string(a) + str[j:]
}

func step(password string, instruction string) string {
	words := strings.Split(instruction, " ")

	switch words[0] {
	case "swap":
		if words[1] == "position" {
			a, _ := strconv.Atoi(words[2])
			b, _ := strconv.Atoi(words[5])
			password = swapPosition(password, a, b)
		} else {
			a, b := words[2], words[5]
			password = swapLetter(password, a, b)
		}
	case "rotate":
		if words[1] == "based" {
			letter := words[6]
			password = rotateBased(password, letter)
		} else {
			direction := words[1]
			steps, _ := strconv.Atoi(words[2])
			password = rotate(password, direction, steps)
		}
	case "reverse":
		a, _ := strconv.Atoi(words[2])
		b, _ := strconv.Atoi(words[4])
		password = reverse(password, a, b)
	case "move":
		a, _ := strconv.Atoi(words[2])
		b, _ := strconv.Atoi(words[5])
		password = move(password, a, b)
	}

	return password
}

func solve1(start string, instructions []string) string {
	password := start

	for _, instruction := range instructions {
		password = step(password, instruction)
	}

	return password
}

func permutations(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	var result []string

	for i, c := range str {
		for _, p := range permutations(str[:i] + str[i+1:]) {
			result = append(result, string(c)+p)
		}
	}

	return result
}

func solve2(start string, instructions []string) string {
	password := start

	for _, p := range permutations(password) {
		if solve1(p, instructions) == password {
			return p
		}
	}

	return ""
}

func Solve(input string) string {
	instructions := strings.Split(strings.TrimSpace(input), "\n")

	return "Day21\nPart1: " + solve1("abcdefgh", instructions) + "\nPart2: " + solve2("fbgdceah", instructions)
}
