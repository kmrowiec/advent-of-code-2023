package day21

import (
	"fmt"

	"dev.kmrowiec/aoc/helper/grid"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day21/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day21/input/input1.txt")
}

func Simulate(inputFile string, steps int) int {
	g := grid.GridFromFile(inputFile)
	startingPosition := g.FindLocation('S')

	currentPoints := map[grid.Point]bool{startingPosition: true}
	for i := 0; i < steps; i++ {
		newPoints := map[grid.Point]bool{}
		for p := range currentPoints {
			for _, a := range g.GetAdjacents(p, false) {
				if g.GetCharAt(a) != '#' {
					newPoints[a] = true
				}
			}
		}
		fmt.Printf("Step %d: %v\n", i, newPoints)
		currentPoints = newPoints
	}

	fmt.Println("Starting location is ", startingPosition)
	return len(currentPoints)
}

func SolvePart1(inputFile string) string {
	return fmt.Sprint(Simulate(inputFile, 64))
}

func SolvePart2(inputFile string) string {
	result := 0
	return fmt.Sprint(result)
}
