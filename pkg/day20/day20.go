package day20

import (
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

func (s IntervalSlice) Len() int {
	return len(s)
}

func (s IntervalSlice) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func (s IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func solve1(intervals IntervalSlice) int {
	sort.Sort(intervals)

	min := 0
	for _, interval := range intervals {
		if interval.Start > min {
			return min
		}
		if interval.End > min {
			min = interval.End + 1
		}
	}

	return min
}

func solve2(intervals IntervalSlice) int {
	sort.Sort(intervals)

	mn, mx := intervals[0].Start, intervals[0].End
	total := 0

	for _, interval := range intervals {
		if interval.Start > mx+1 {
			total += mx - mn + 1
			mn, mx = interval.Start, interval.End
		} else {
			if interval.End > mx {
				mx = interval.End
			}
		}
	}

	return 4294967296 - total - (mx - mn + 1)
}

func Solve(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	intervals := make([]Interval, 0)
	for _, line := range lines {
		points := strings.Split(line, "-")
		start, _ := strconv.Atoi(points[0])
		end, _ := strconv.Atoi(points[1])
		intervals = append(intervals, Interval{Start: start, End: end})
	}

	return "Day20\nPart1: " + strconv.Itoa(solve1(intervals)) + "\nPart2: " + strconv.Itoa(solve2(intervals))
}
