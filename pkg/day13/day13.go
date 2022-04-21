package day13

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func isWall(p Point, designerNumber int) bool {
	x := p.X
	y := p.Y

	sum := x*x + 3*x + 2*x*y + y + y*y + designerNumber

	ones := 0
	for sum > 0 {
		if sum&1 == 1 {
			ones++
		}
		sum >>= 1
	}

	return ones%2 == 1
}

func isInBounds(p Point) bool {
	return p.X >= 0 && p.Y >= 0
}

func bfs(start Point, end Point, designedNumber int) []Point {
	queue := []Point{start}
	visited := make(map[Point]bool)
	parents := make(map[Point]Point)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current] = true

		if current == end {
			break
		}

		for _, direction := range []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			next := Point{current.X + direction.X, current.Y + direction.Y}

			if !isWall(next, designedNumber) && !visited[next] {
				parents[next] = current
				queue = append(queue, next)
			}
		}
	}

	if !visited[end] {
		return []Point{}
	}

	path := []Point{end}
	for current := parents[end]; current != start; current = parents[current] {
		path = append([]Point{current}, path...)
	}

	return path
}

func solve1(start Point, end Point, designedNumber int) int {
	return len(bfs(start, end, designedNumber))
}

func bfs2(start Point, designedNumber int) []Point {
	queue := []Point{start}
	visited := make(map[Point]int)
	parents := make(map[Point]Point)

	visited[start] = 0
	parents[start] = start

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visited[current] = 1 + visited[parents[current]]

		if visited[current] > 50 {
			continue
		}

		for _, direction := range []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			next := Point{current.X + direction.X, current.Y + direction.Y}

			_, vis := visited[next]
			if !isWall(next, designedNumber) && isInBounds(next) && !vis {
				parents[next] = current
				queue = append(queue, next)
			}
		}
	}

	points := []Point{}
	for p := range visited {
		points = append(points, p)
	}

	return points
}

func solve2(start Point, designedNumber int) int {
	return len(bfs2(start, designedNumber))
}

func Solve(input string) string {
	designerNumber, _ := strconv.Atoi(strings.TrimSpace(input))

	return "Day13\nPart1: " + strconv.Itoa(solve1(Point{1, 1}, Point{31, 39}, designerNumber)) + "\nPart2: " + strconv.Itoa(solve2(Point{1, 1}, designerNumber))
}
