package day04

import (
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	Name     string
	SectorId int
	Checksum string
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func ParseRoom(room string) Room {
	parts := strings.Split(room, "-")

	last := parts[len(parts)-1]
	init := strings.Join(parts[:len(parts)-1], "")

	lastParts := strings.Split(last, "[")
	sectorId, _ := strconv.Atoi(lastParts[0])
	checksum := lastParts[1][:len(lastParts[1])-1]

	return Room{init, sectorId, checksum}
}

func IsValid(room Room) bool {
	letters := map[string]int{}

	for _, letter := range room.Name {
		letters[string(letter)]++
	}

	pairs := make(PairList, len(letters))

	i := 0
	for k, v := range letters {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pairs))

	keys := []string{}

	for i, pair := range pairs {
		keys = append(keys, pair.Key)

		if i == 4 {
			break
		}
	}

	return room.Checksum == strings.Join(keys, "")
}

func Solve1(rooms []Room) int {
	var sum int
	for _, room := range rooms {
		if IsValid(room) {
			sum += room.SectorId
		}
	}

	return sum
}

func Solve2(rooms []Room) int {
	for _, room := range rooms {
		n := room.SectorId % 26
		name := ""

		for _, letter := range room.Name {
			letter = letter + rune(n)
			if letter > 'z' {
				letter = letter - 26
			}
			name += string(letter)
		}

		if name == "northpoleobjectstorage" {
			return room.SectorId
		}
	}

	return 0
}

func Solve(input string) string {
	rooms := []Room{}

	for _, room := range strings.Split(strings.TrimSpace(input), "\n") {
		rooms = append(rooms, ParseRoom(room))
	}

	return "Day04\nPart1: " + strconv.Itoa(Solve1(rooms)) + "\nPart2: " + strconv.Itoa(Solve2(rooms))
}
