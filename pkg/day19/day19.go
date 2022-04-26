package day19

import (
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

func solve1(n int) int {
	elf := &Node{Value: 1}
	it := elf
	for i := 2; i <= n; i++ {
		it.Next = &Node{Value: i}
		it = it.Next
	}
	it.Next = elf

	for elf.Next != elf {
		elf.Next = elf.Next.Next
		elf = elf.Next
	}

	return elf.Value
}

func solve2(n int) int {
	i := 1

	for i*3 < n {
		i *= 3
	}

	return (n - i)
}

func Solve(input string) string {
	n, _ := strconv.Atoi(strings.TrimSpace(input))

	return "Day19\nPart1: " + strconv.Itoa(solve1(n)) + "\nPart2: " + strconv.Itoa(solve2(n))
}
