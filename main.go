package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/alexjercan/aoc-2016/pkg/day01"
	"github.com/alexjercan/aoc-2016/pkg/day02"
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
	}

	panic("Unknown day")
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
