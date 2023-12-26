package main

import (
	"fmt"

	"dev.kmrowiec/aoc/day1"
	"dev.kmrowiec/aoc/day10"
	"dev.kmrowiec/aoc/day11"
	"dev.kmrowiec/aoc/day12"
	"dev.kmrowiec/aoc/day13"
	"dev.kmrowiec/aoc/day14"
	"dev.kmrowiec/aoc/day15"
	"dev.kmrowiec/aoc/day16"
	"dev.kmrowiec/aoc/day17"
	"dev.kmrowiec/aoc/day18"
	"dev.kmrowiec/aoc/day19"
	"dev.kmrowiec/aoc/day2"
	"dev.kmrowiec/aoc/day3"
	"dev.kmrowiec/aoc/day4"
	"dev.kmrowiec/aoc/day5"
	"dev.kmrowiec/aoc/day6"
	"dev.kmrowiec/aoc/day7"
	"dev.kmrowiec/aoc/day8"
	"dev.kmrowiec/aoc/day9"
	"dev.kmrowiec/aoc/helper"
)

func main() {
	solvers := map[int]helper.Solver{
		1:  &day1.Solver{},
		2:  &day2.Solver{},
		3:  &day3.Solver{},
		4:  &day4.Solver{},
		5:  &day5.Solver{},
		6:  &day6.Solver{},
		7:  &day7.Solver{},
		8:  &day8.Solver{},
		9:  &day9.Solver{},
		10: &day10.Solver{},
		11: &day11.Solver{},
		12: &day12.Solver{},
		13: &day13.Solver{},
		14: &day14.Solver{},
		15: &day15.Solver{},
		16: &day16.Solver{},
		17: &day17.Solver{},
		18: &day18.Solver{},
		19: &day19.Solver{},
	}

	day := 19
	// fmt.Printf("Day %d, Part 1: %v \n", day, solvers[day].PartOne())
	fmt.Printf("Day %d, Part 2: %v \n", day, solvers[day].PartTwo())
}
