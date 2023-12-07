package main

import (
	"fmt"

	"dev.kmrowiec/aoc/day1"
	"dev.kmrowiec/aoc/day2"
	"dev.kmrowiec/aoc/day3"
	"dev.kmrowiec/aoc/day4"
	"dev.kmrowiec/aoc/day5"
	"dev.kmrowiec/aoc/day6"
	"dev.kmrowiec/aoc/day7"
	"dev.kmrowiec/aoc/helper"
)

func main() {
	solvers := map[int]helper.Solver{
		1: &day1.Solver{},
		2: &day2.Solver{},
		3: &day3.Solver{},
		4: &day4.Solver{},
		5: &day5.Solver{},
		6: &day6.Solver{},
		7: &day7.Solver{},
	}

	day := 7
	fmt.Printf("Day %d, Part 1: %v \n", day, solvers[day].PartOne())
	fmt.Printf("Day %d, Part 2: %v \n", day, solvers[day].PartTwo())
}
