package day14

import (
	"fmt"

	"dev.kmrowiec/aoc/helper/grid"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day14/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day14/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	result := 0
	g := grid.GridFromFile(inputFile)

	maxLoad := g.ColumnLength()
	for x := 0; x < g.RowLength(); x++ {
		spaces := 0
		column := g.GetColumn(x)
		for index, char := range column {
			switch char {
			case '.':
				spaces++
			case '#':
				spaces = 0
			case 'O':
				score := maxLoad - index + spaces
				result += score
			}
		}
	}

	return fmt.Sprint(result)
}

func spin(g *grid.Grid) {
	// Tilt north
	for x := 0; x < g.RowLength(); x++ {
		spaces := 0
		column := g.GetColumn(x)
		for index, char := range column {
			switch char {
			case '.':
				spaces++
			case '#':
				spaces = 0
				g.SetCharAt(grid.PointFromXY(x, index), '#')
			case 'O':
				g.SetCharAt(grid.PointFromXY(x, index), '.')
				g.SetCharAt(grid.PointFromXY(x, index-spaces), 'O')
			}
		}
	}

	// Tilt west
	for y := 0; y < g.ColumnLength(); y++ {
		spaces := 0
		row := g.GetRow(y)
		for index := 0; index < g.RowLength(); index++ {
			char := row[index]
			switch char {
			case '.':
				spaces++
			case '#':
				spaces = 0
				g.SetCharAt(grid.PointFromXY(index, y), '#')
			case 'O':
				g.SetCharAt(grid.PointFromXY(index, y), '.')
				g.SetCharAt(grid.PointFromXY(index-spaces, y), 'O')
			}
		}
	}

	// Tilt south
	for x := 0; x < g.RowLength(); x++ {
		spaces := 0
		column := g.GetColumn(x)
		for index := g.ColumnLength() - 1; index >= 0; index-- {
			char := column[index]
			switch char {
			case '.':
				spaces++
			case '#':
				spaces = 0
				g.SetCharAt(grid.PointFromXY(x, index), '#')
			case 'O':
				g.SetCharAt(grid.PointFromXY(x, index), '.')
				g.SetCharAt(grid.PointFromXY(x, index+spaces), 'O')
			}
		}
	}

	// Tilt east
	for y := 0; y < g.ColumnLength(); y++ {
		spaces := 0
		row := g.GetRow(y)
		for index := g.RowLength() - 1; index >= 0; index-- {
			char := row[index]
			switch char {
			case '.':
				spaces++
			case '#':
				spaces = 0
				g.SetCharAt(grid.PointFromXY(index, y), '#')
			case 'O':
				g.SetCharAt(grid.PointFromXY(index, y), '.')
				g.SetCharAt(grid.PointFromXY(index+spaces, y), 'O')
			}
		}
	}
}

func calculateLoad(g *grid.Grid) int {
	result := 0
	maxLoad := g.ColumnLength()
	for x := 0; x < g.RowLength(); x++ {
		for index, char := range g.GetColumn(x) {
			if char == 'O' {
				result += maxLoad - index
			}
		}
	}
	return result
}

func SolvePart2(inputFile string) string {
	g := grid.GridFromFile(inputFile)

	// The spinning will eventually settle into a repeating pattern
	// If we find out the frequency and starting index of the repeating pattern,
	// we can calculate the total load at 1 billion spins without actually needing to simulate them

	const spins = 1000 // just a number large enough to notice a repeating pattern
	for i := 0; i < spins; i++ {
		spin(&g)
		fmt.Printf("Reached %d after %d spins \n", calculateLoad(&g), i)
	}
	return fmt.Sprint(0) // The rest of the answer can simply be manual
}
