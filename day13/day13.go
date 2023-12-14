package day13

import (
	"fmt"

	"dev.kmrowiec/aoc/helper"
	"dev.kmrowiec/aoc/helper/grid"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day13/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day13/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	result := 0
	verticalCount, horizontalCount := 0, 0
	currentLines := make([]string, 0)
	for _, line := range helper.ReadInputFile(inputFile) {
		if line == "" {
			grid := grid.GridFromLines(currentLines)
			vertical, count := Reflect(&grid)
			if vertical {
				verticalCount += count
			} else {
				horizontalCount += count
			}
			currentLines = make([]string, 0)
		} else {
			currentLines = append(currentLines, line)
		}
	}

	grid := grid.GridFromLines(currentLines)
	vertical, count := Reflect(&grid)
	if vertical {
		verticalCount += count
	} else {
		horizontalCount += count
	}

	result = verticalCount + (100 * horizontalCount)

	return fmt.Sprint(result)
}

func Reflect(grid *grid.Grid) (vertical bool, reflectedRows int) {

	fmt.Println("Finding reflection in grid:")
	grid.Draw()

	//First scan for horizontal reflection
	for y := 0; y < grid.ColumnLength()-1; y++ {

		if grid.Rows[y] == grid.Rows[y+1] {

			fmt.Printf("Line %d and %d match! \n", y, y+1)
			reflecting := true
			for i := 1; y-i >= 0 && y+1+i < grid.ColumnLength(); i++ {
				fmt.Print("Rows " + grid.Rows[y-i] + " and " + grid.Rows[y+1+i])
				if grid.Rows[y-i] != grid.Rows[y+1+i] {
					reflecting = false
					fmt.Println(" ...are different!")
					break
				}
				fmt.Println(" ...match!")
			}
			if reflecting {
				fmt.Printf("Found horizontal reflection with %d rows above it \n", y+1)
				return false, y + 1
			}

		}
	}

	//Then scan for vertical reflection
	for x := 0; x < grid.RowLength()-1; x++ {
		if grid.GetColumn(x) == grid.GetColumn(x+1) {

			fmt.Printf("Columns %d and %d match! \n", x, x+1)
			reflecting := true
			for i := 1; x-i >= 0 && x+1+i < grid.RowLength(); i++ {
				fmt.Printf("Columns (%d) %v and (%d) %v", x-i, grid.GetColumn(x-i), x+1+i, grid.GetColumn(x+1+i))
				if grid.GetColumn(x-i) != grid.GetColumn(x+1+i) {
					reflecting = false
					fmt.Println(" ...are different!")
					break
				}
				fmt.Println(" ...match!")
			}
			if reflecting {
				fmt.Printf("Found horizontal reflection with %d columns left to it \n", x)
				return true, x + 1
			}

		}
	}
	return false, -1
}

func ReflectWithOneSmudge(grid *grid.Grid) (vertical bool, reflectedRows int) {

	smudgeFound := false

	//First scan for horizontal reflection
	for y := 0; y < grid.ColumnLength()-1; y++ {

		if grid.Rows[y] == grid.Rows[y+1] {

			fmt.Printf("Line %d and %d match! \n", y, y+1)
			reflecting := true
			for i := 1; y-i >= 0 && y+1+i < grid.ColumnLength(); i++ {
				fmt.Print("Rows " + grid.Rows[y-i] + " and " + grid.Rows[y+1+i])
				if grid.Rows[y-i] != grid.Rows[y+1+i] {
					if smudgeFound {
						reflecting = false
						fmt.Println(" ...are different!")
						break
					} else {
						smudgeFound = true
					}
				}
				fmt.Println(" ...match!")
			}
			if reflecting && smudgeFound {
				fmt.Printf("Found horizontal reflection with %d rows above it \n", y+1)
				return false, y + 1
			}

		}
	}

	smudgeFound = false

	//Then scan for vertical reflection
	for x := 0; x < grid.RowLength()-1; x++ {
		if grid.GetColumn(x) == grid.GetColumn(x+1) {

			fmt.Printf("Columns %d and %d match! \n", x, x+1)
			reflecting := true
			for i := 1; x-i >= 0 && x+1+i < grid.RowLength(); i++ {
				fmt.Printf("Columns (%d) %v and (%d) %v", x-i, grid.GetColumn(x-i), x+1+i, grid.GetColumn(x+1+i))
				if grid.GetColumn(x-i) != grid.GetColumn(x+1+i) {
					if smudgeFound {
						reflecting = false
						fmt.Println(" ...are different!")
						break
					} else {
						smudgeFound = true
					}
				}
				fmt.Println(" ...match!")
			}
			if reflecting && smudgeFound {
				fmt.Printf("Found horizontal reflection with %d columns left to it \n", x)
				return true, x + 1
			}

		}
	}
	return false, -1
}

func isDifferentByOne(first, second string) bool {
	differences := 0
	for k := range first {
		if first[k] != second[k] {
			differences++
		}
	}
	return differences == 1
}

func FindVariations(g grid.Grid) []grid.Grid {
	fmt.Println("Finding variations.")

	result := make([]grid.Grid, 0)

	for y := 0; y < g.ColumnLength()-1; y++ {
		if index, ok := findSmudge(g.Rows[y], g.Rows[y+1]); ok {
			newGrid := grid.GridFromLines(g.Rows)
			newRune := '#'
			if g.GetCharAtXY(index, y) == '#' {
				newRune = '.'
			}
			newGrid.SetCharAt(grid.PointFromXY(index, y), newRune)
			result = append(result, newGrid)
		}
	}

	for x := 0; x < g.RowLength()-1; x++ {
		if index, ok := findSmudge(g.GetColumn(x), g.GetColumn(x+1)); ok {
			newGrid := grid.GridFromLines(g.Rows)
			newRune := '#'
			if g.GetCharAtXY(x, index) == '#' {
				newRune = '.'
			}
			newGrid.SetCharAt(grid.PointFromXY(x, index), newRune)
			result = append(result, newGrid)
		}

	}
	fmt.Printf("Found %d variations.\n", len(result))
	return result
}

func findSmudge(first, second string) (int, bool) {
	// fmt.Printf("Finding a smudge in %v (%d) and %v (%d)\n", first, len(first), second, len(second))
	if isDifferentByOne(first, second) {
		for i := 0; i < len(first); i++ {
			if first[i] != second[i] {
				return i, true
			}
		}
	}
	return 0, false
}

func LoadGridsFromFile(inputFile string) []grid.Grid {
	inputGrids := make([]grid.Grid, 0)
	currentLines := make([]string, 0)
	for _, line := range helper.ReadInputFile(inputFile) {
		if line == "" {
			inputGrids = append(inputGrids, grid.GridFromLines(currentLines))
			currentLines = make([]string, 0)
		} else {
			currentLines = append(currentLines, line)
		}
	}
	inputGrids = append(inputGrids, grid.GridFromLines(currentLines))
	return inputGrids
}

func SolvePart2(inputFile string) string {
	result := 0
	verticalCount, horizontalCount := 0, 0

	inputGrids := LoadGridsFromFile(inputFile)

	fmt.Printf("Loaded %d grids.\n", len(inputGrids))

	for _, g := range inputGrids {
		fmt.Println("")
		g.Draw()

		vertical, count := ReflectWithOneSmudge(&g)
		if count != -1 {
			if vertical {
				verticalCount += count
			} else {
				horizontalCount += count
			}
		} else {
			found := false
			for _, variation := range FindVariations(g) {
				vertical, count := Reflect(&variation)
				if count != -1 {
					if vertical {
						verticalCount += count
					} else {
						horizontalCount += count
					}
					found = true
					break
				}
			}
			if !found {
				panic("did not find a solution at all!")
			}
		}
	}

	result = verticalCount + (100 * horizontalCount)

	return fmt.Sprint(result)
}
