package day24

import (
	"container/heap"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func isValid(point Point, ventMap []string) bool {
	if point.X < 0 || point.Y < 0 || point.X >= len(ventMap) || point.Y >= len(ventMap[0]) {
		return false
	}

	return ventMap[point.X][point.Y] != '#'
}

func bfs(start Point, end Point, ventMap []string) int {
	queue := []Point{start}
	visited := map[Point]bool{start: true}
	parent := map[Point]Point{start: start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range []Point{
			{current.X, current.Y - 1},
			{current.X, current.Y + 1},
			{current.X - 1, current.Y},
			{current.X + 1, current.Y},
		} {
			if _, ok := visited[neighbor]; !ok && isValid(neighbor, ventMap) {
				visited[neighbor] = true
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	if _, ok := visited[end]; !ok {
		return -1
	}

	path := []Point{end}
	for current := end; current != start; current = parent[current] {
		path = append(path, current)
	}

	return len(path) - 1
}

func createWeightedGraph(ventMap []string) map[Point]map[Point]int {
	graph := make(map[Point]map[Point]int)

	positions := make([]Point, 0)
	for row := 0; row < len(ventMap); row++ {
		for col := 0; col < len(ventMap[row]); col++ {
			if string(ventMap[row][col]) >= "0" && string(ventMap[row][col]) <= "9" {
				positions = append(positions, Point{row, col})
			}
		}
	}

	for _, p := range positions {
		graph[p] = make(map[Point]int)
	}

	for i, start := range positions {
		for j, end := range positions {
			if i >= j {
				continue
			}

			weight := bfs(start, end, ventMap)
			graph[start][end] = weight
			graph[end][start] = weight
		}
	}

	return graph
}

type Item struct {
	value    int
	priority int
	mask     int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, mask int, priority int) {
	item.value = value
	item.mask = mask
	item.priority = priority
	heap.Fix(pq, item.index)
}

func solve1(graph [][]int) int {
	cost := make([]map[int]int, len(graph))
	pq := make(PriorityQueue, len(graph))

	for i := range graph {
		cost[i] = make(map[int]int)
		cost[i][1<<i] = 0
		pq[i] = &Item{
			value:    i,
			mask:     1 << i,
			priority: 1000000,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)

		for neighbor := range graph[item.value] {
			mask := item.mask | (1 << neighbor)
			priority := cost[item.value][item.mask] + graph[item.value][neighbor]

			if oldPriority, ok := cost[neighbor][mask]; !ok || oldPriority > priority {
				heap.Push(&pq, &Item{
					value:    neighbor,
					priority: priority,
					mask:     mask,
				})
				cost[neighbor][mask] = priority
			}
		}
	}

	mask := (1 << len(graph)) - 1
	minCost := 0
	for i := range graph {
		if cost[i][mask] < minCost || minCost == 0 {
			minCost = cost[i][mask]
		}
	}

	return minCost
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	graphMap := createWeightedGraph(lines)

	nodes := make([]Point, 0)
	for node := range graphMap {
		nodes = append(nodes, node)
	}

	graph := make([][]int, len(nodes))
	for i := range graph {
		graph[i] = make([]int, len(nodes))
		for j := range graph[i] {
			graph[i][j] = graphMap[nodes[i]][nodes[j]]
		}
	}

	return "Day24\nPart1: " + strconv.Itoa(solve1(graph)) + "\nPart2: " + "0"
}
