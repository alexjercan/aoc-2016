package day22

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Used  int
	Avail int
}

type Position struct {
	Row int
	Col int
}

func solve1(nodes []Node) int {
	count := 0

	for i := range nodes {
		if nodes[i].Used == 0 {
			continue
		}

		for j := range nodes {
			if i == j {
				continue
			}

			if nodes[j].Avail >= nodes[i].Used {
				count++
			}
		}
	}

	return count
}

func getNeighbors(position Position, rows int, cols int) []Position {
	neighbors := make([]Position, 0)

	if position.Row > 0 {
		neighbors = append(neighbors, Position{position.Row - 1, position.Col})
	}
	if position.Row < rows-1 {
		neighbors = append(neighbors, Position{position.Row + 1, position.Col})
	}
	if position.Col > 0 {
		neighbors = append(neighbors, Position{position.Row, position.Col - 1})
	}
	if position.Col < cols-1 {
		neighbors = append(neighbors, Position{position.Row, position.Col + 1})
	}

	return neighbors
}

func remove(slice []Position, s Position) []Position {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func solve2(nodes []Node) (int, string) {
	positions := make([]Position, len(nodes))
	cols, rows := 0, 0
	for i := range nodes {
		words := strings.Split(nodes[i].Name, "-")
		col, _ := strconv.Atoi(words[1][1:])
		row, _ := strconv.Atoi(words[2][1:])
		positions[i] = Position{row, col}

		if col > cols {
			cols = col
		}
		if row > rows {
			rows = row
		}
	}

	nodesMap := make([][]Node, rows+1)
	for i := range nodesMap {
		nodesMap[i] = make([]Node, cols+1)
	}
	for i, node := range nodes {
		nodesMap[positions[i].Row][positions[i].Col] = node
	}

	str := ""
	holeRow, holeCol := 0, 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if nodesMap[i][j].Used == 0 {
				str += "_"
				holeRow, holeCol = i, j
			} else if nodesMap[i][j].Used >= 100 {
				str += "#"
			} else {
				str += "."
			}
		}

		str += "\n"
	}

	fmt.Println(holeRow, holeCol)

	steps1 := holeRow + holeCol + cols - 1
	steps2 := 5*(cols-1) + 1

	return steps1 + steps2, str
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")[2:]
	nodes := make([]Node, len(lines))
	for i, line := range lines {
		words := strings.Fields(line)
		used, _ := strconv.Atoi(words[2][:len(words[2])-1])
		avail, _ := strconv.Atoi(words[3][:len(words[3])-1])
		nodes[i] = Node{
			Name:  words[0],
			Used:  used,
			Avail: avail,
		}
	}

	steps, draw := solve2(nodes)
	return "Day22\nPart1: " + strconv.Itoa(solve1(nodes)) + "\nPart2: " + strconv.Itoa(steps) + "\n" + draw
}
