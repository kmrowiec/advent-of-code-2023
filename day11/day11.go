package day11

import (
	"fmt"

	"dev.kmrowiec/aoc/helper"
	"dev.kmrowiec/aoc/helper/grid"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day11/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day11/input/input1.txt")
}

type Galaxy struct {
	Code  int
	Point grid.Point
}

func FindGalaxies(universe *grid.Grid) []Galaxy {
	result := make([]Galaxy, 0)
	galaxyId := 1
	for y := 0; y < universe.ColumnLength(); y++ {
		for x := 0; x < universe.RowLength(); x++ {
			if universe.GetCharAtXY(x, y) == '#' {
				result = append(result, Galaxy{Code: galaxyId, Point: grid.PointFromXY(x, y)})
				galaxyId++
			}
		}
	}
	return result
}

func FindPairsOfGalaxies(galaxies []Galaxy) [][]Galaxy {
	result := make([][]Galaxy, 0)
	startingB := 0
	for a := 0; a < len(galaxies); a++ {
		for b := startingB; b < len(galaxies); b++ {
			if a != b {
				result = append(result, []Galaxy{galaxies[a], galaxies[b]})
			}
		}
		startingB++
	}
	return result
}

func expandUniverse(grid *grid.Grid) (map[int]bool, map[int]bool) {

	expansionColumns, expansionRows := map[int]bool{}, map[int]bool{}

	for rowNumber := 0; rowNumber < grid.ColumnLength(); rowNumber++ {
		empty := true
		for _, char := range grid.Rows[rowNumber] {
			if char != '.' {
				empty = false
				break
			}
		}
		if empty {
			expansionRows[rowNumber] = true
		}
	}

	for columnNumber := 0; columnNumber < grid.RowLength(); columnNumber++ {
		empty := true
		for y := range grid.Rows {
			if grid.GetCharAtXY(columnNumber, y) != '.' {
				empty = false
				break
			}
		}
		if empty {
			expansionColumns[columnNumber] = true
		}
	}

	return expansionColumns, expansionRows
}

func SumOfShortestPaths(inputFile string, multiplier int) string {
	result := 0
	grid := grid.GridFromFile(inputFile)
	galaxies := FindGalaxies(&grid)
	pairs := FindPairsOfGalaxies(galaxies)
	expandedColumns, expandedRows := expandUniverse(&grid)

	for _, pair := range pairs {
		result += pair[0].Point.DistanceTo(pair[1].Point)

		lowerX, higherX := helper.MinAndMax(pair[0].Point.X, pair[1].Point.X)
		for x := lowerX; x < higherX; x++ {
			if _, ok := expandedColumns[x]; ok {
				result += multiplier - 1
			}
		}

		lowerY, higherY := helper.MinAndMax(pair[0].Point.Y, pair[1].Point.Y)
		for y := lowerY; y < higherY; y++ {
			if _, ok := expandedRows[y]; ok {
				result += multiplier - 1
			}
		}
	}

	return fmt.Sprint(result)
}

func SolvePart1(inputFile string) string {
	return SumOfShortestPaths(inputFile, 2)
}

func SolvePart2(inputFile string) string {
	return SumOfShortestPaths(inputFile, 1000000)
}
