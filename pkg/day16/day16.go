package day16

import (
	"strings"
)

func dragonCurveStep(a []byte) []byte {
	b := make([]byte, 0)
	b = append(b, a...)
	b = append(b, 0)

	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, 1-a[i])
	}

	return b
}

func dragonCurve(a []byte, length int) []byte {
	for len(a) < length {
		a = dragonCurveStep(a)
	}

	return a[:length]
}

func checksumStep(a []byte) []byte {
	b := make([]byte, 0)

	for i := 0; i < len(a); i += 2 {
		if a[i] == a[i+1] {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}

	return b
}

func checksum(a []byte) []byte {
	for len(a)%2 == 0 {
		a = checksumStep(a)
	}

	return a
}

func solve1(a []byte) []byte {
	return checksum(dragonCurve(a, 272))
}

func solve2(a []byte) []byte {
	return checksum(dragonCurve(a, 35651584))
}

func prepareInput(input string) []byte {
	a := make([]byte, 0)

	for _, c := range input {
		if c == '1' {
			a = append(a, 1)
		} else {
			a = append(a, 0)
		}
	}

	return a
}

func prepareOutput(output []byte) string {
	a := make([]byte, 0)

	for _, c := range output {
		if c == 1 {
			a = append(a, '1')
		} else {
			a = append(a, '0')
		}
	}

	return string(a)
}

func Solve(input string) string {
	a := prepareInput(strings.TrimSpace(input))

	return "Day16\nPart1: " + prepareOutput(solve1(a)) + "\nPart2: " + prepareOutput(solve2(a))
}
