package day03

import (
	"strconv"
	"strings"
)

type Lengths struct {
	L1 int
	L2 int
	L3 int
}

func Solve1(lengths []Lengths) int {
	triangles := 0

	for _, l := range lengths {
		if l.L1+l.L2 > l.L3 && l.L2+l.L3 > l.L1 && l.L1+l.L3 > l.L2 {
			triangles++
		}
	}

	return triangles
}

func Solve(input string) string {
	lengths := []Lengths{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)

		l1, _ := strconv.Atoi(fields[0])
		l2, _ := strconv.Atoi(fields[1])
		l3, _ := strconv.Atoi(fields[2])

		lengths = append(lengths, Lengths{l1, l2, l3})
	}

	lengths2 := []Lengths{}
	numbers1 := []int{}
	numbers2 := []int{}
	numbers3 := []int{}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)

		l1, _ := strconv.Atoi(fields[0])
		l2, _ := strconv.Atoi(fields[1])
		l3, _ := strconv.Atoi(fields[2])

		numbers1 = append(numbers1, l1)
		numbers2 = append(numbers2, l2)
		numbers3 = append(numbers3, l3)
	}

	numbers := append(numbers1, numbers2...)
	numbers = append(numbers, numbers3...)

	for i := 0; i < len(numbers); i += 3 {
		lengths2 = append(lengths2, Lengths{numbers[i], numbers[i+1], numbers[i+2]})
	}

	return "Day03\nPart1: " + strconv.Itoa(Solve1(lengths)) + "\nPart2: " + strconv.Itoa(Solve1(lengths2))
}
