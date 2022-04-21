package day17

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type QueueT struct {
	p Position
	c string
}

func isOpen(hash byte) bool {
	return hash == 'b' || hash == 'c' || hash == 'd' || hash == 'e' || hash == 'f'
}

func openDoors(code string) (bool, bool, bool, bool) {
	hash := md5.Sum([]byte(code))
	hashStr := hex.EncodeToString(hash[:])

	return isOpen(hashStr[0]), isOpen(hashStr[1]), isOpen(hashStr[2]), isOpen(hashStr[3])
}

func isValid(position Position) bool {
	return position.x >= 0 && position.x < 4 && position.y >= 0 && position.y < 4
}

func bfs(end Position, queue []QueueT) (string, bool) {
	if len(queue) == 0 {
		return "", false
	}

	start := queue[0].p
	code := queue[0].c
	queue = queue[1:]

	if start == end {
		return code, true
	}

	up, down, left, right := openDoors(code)

	if up && isValid(Position{start.x, start.y - 1}) {
		queue = append(queue, QueueT{Position{start.x, start.y - 1}, code + "U"})
	}

	if down && isValid(Position{start.x, start.y + 1}) {
		queue = append(queue, QueueT{Position{start.x, start.y + 1}, code + "D"})
	}

	if left && isValid(Position{start.x - 1, start.y}) {
		queue = append(queue, QueueT{Position{start.x - 1, start.y}, code + "L"})
	}

	if right && isValid(Position{start.x + 1, start.y}) {
		queue = append(queue, QueueT{Position{start.x + 1, start.y}, code + "R"})
	}

	return bfs(end, queue)
}

func dfs(start Position, end Position, code string) []string {
	if start == end {
		return []string{code}
	}

	paths := []string{}
	up, down, left, right := openDoors(code)

	if up && isValid(Position{start.x, start.y - 1}) {
		codeU := dfs(Position{start.x, start.y - 1}, end, code+"U")
		paths = append(paths, codeU...)
	}

	if down && isValid(Position{start.x, start.y + 1}) {
		codeD := dfs(Position{start.x, start.y + 1}, end, code+"D")
		paths = append(paths, codeD...)
	}

	if left && isValid(Position{start.x - 1, start.y}) {
		codeL := dfs(Position{start.x - 1, start.y}, end, code+"L")
		paths = append(paths, codeL...)
	}

	if right && isValid(Position{start.x + 1, start.y}) {
		codeR := dfs(Position{start.x + 1, start.y}, end, code+"R")
		paths = append(paths, codeR...)
	}

	return paths
}

func solve1(code string) string {
	output, _ := bfs(Position{3, 3}, []QueueT{{Position{0, 0}, code}})

	return output[len(code):]
}

func solve2(code string) int {
	paths := dfs(Position{0, 0}, Position{3, 3}, code)

	maxlen := 0
	for _, path := range paths {
		path = path[len(code):]
		if len(path) > maxlen {
			maxlen = len(path)
		}
	}

	return maxlen
}

func Solve(input string) string {
	code := strings.TrimSpace(input)

	return "Day17\nPart1: " + solve1(code) + "\nPart2: " + strconv.Itoa(solve2(code))
}
