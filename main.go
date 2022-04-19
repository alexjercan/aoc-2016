package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/alexjercan/aoc-2016/pkg/day01"
	"github.com/alexjercan/aoc-2016/pkg/day02"
	"github.com/alexjercan/aoc-2016/pkg/day03"
	"github.com/alexjercan/aoc-2016/pkg/day04"
	"github.com/alexjercan/aoc-2016/pkg/day05"
	"github.com/alexjercan/aoc-2016/pkg/day06"
	"github.com/alexjercan/aoc-2016/pkg/day07"
	"github.com/alexjercan/aoc-2016/pkg/day08"
	"github.com/alexjercan/aoc-2016/pkg/day09"
	"github.com/alexjercan/aoc-2016/pkg/day10"
	"github.com/alexjercan/aoc-2016/pkg/day11"
	"github.com/alexjercan/aoc-2016/pkg/day12"
	"github.com/alexjercan/aoc-2016/pkg/day13"
	"github.com/alexjercan/aoc-2016/pkg/day14"
	"github.com/alexjercan/aoc-2016/pkg/day15"
	"github.com/alexjercan/aoc-2016/pkg/day16"
	"github.com/alexjercan/aoc-2016/pkg/day17"
	"github.com/alexjercan/aoc-2016/pkg/day18"
	"github.com/alexjercan/aoc-2016/pkg/day19"
	"github.com/alexjercan/aoc-2016/pkg/day20"
	"github.com/alexjercan/aoc-2016/pkg/day21"
	"github.com/alexjercan/aoc-2016/pkg/day22"
	"github.com/alexjercan/aoc-2016/pkg/day23"
	"github.com/alexjercan/aoc-2016/pkg/day24"
	"github.com/alexjercan/aoc-2016/pkg/day25"
)

func ReadInput(day int) string {
	filename := fmt.Sprintf("input/day%02d.input", day)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func GetSolver(day int) func(string) string {
	switch day {
	case 1:
		return day01.Solve
	case 2:
		return day02.Solve
	case 3:
		return day03.Solve
	case 4:
		return day04.Solve
	case 5:
		return day05.Solve
	case 6:
		return day06.Solve
	case 7:
		return day07.Solve
	case 8:
		return day08.Solve
	case 9:
		return day09.Solve
	case 10:
		return day10.Solve
	case 11:
		return day11.Solve
	case 12:
		return day12.Solve
	case 13:
		return day13.Solve
	case 14:
		return day14.Solve
	case 15:
		return day15.Solve
	case 16:
		return day16.Solve
	case 17:
		return day17.Solve
	case 18:
		return day18.Solve
	case 19:
		return day19.Solve
	case 20:
		return day20.Solve
	case 21:
		return day21.Solve
	case 22:
		return day22.Solve
	case 23:
		return day23.Solve
	case 24:
		return day24.Solve
	case 25:
		return day25.Solve
	default:
		return func(string) string { return "Unknown day" }
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: aoc-2016 day")
		return
	}

	day, _ := strconv.Atoi(os.Args[1])

	input := ReadInput(day)

	solver := GetSolver(day)

	result := solver(input)

	fmt.Println(result)
}
