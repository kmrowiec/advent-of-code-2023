package grid

import (
	"fmt"
	"math"

	"dev.kmrowiec/aoc/helper"
)

type Point struct {
	X, Y int
}

func PointFromXY(x, y int) Point {
	return Point{X: x, Y: y}
}

type Grid struct {
	Rows []string
}

func GridFromFile(inputFile string) Grid {
	lines := helper.ReadInputFile(inputFile)
	return GridFromLines(lines)
}

func GridFromLines(lines []string) Grid {
	grid := Grid{Rows: lines}
	return grid
}

func (grid *Grid) GetCharAt(point Point) rune {
	return rune(grid.Rows[point.Y][point.X])
}

func (grid *Grid) GetCharAtXY(x, y int) rune {
	return grid.GetCharAt(Point{X: x, Y: y})
}

func (grid *Grid) SetCharAt(point Point, char rune) {
	out := []rune(grid.Rows[point.Y])
	out[point.X] = char
	grid.Rows[point.Y] = string(out)
}

func (grid *Grid) FindLocation(content rune) Point {
	for y, row := range grid.Rows {
		for x, char := range row {
			if char == content {
				return Point{X: x, Y: y}
			}
		}
	}
	return Point{}
}

func (grid *Grid) GetRow(y int) string {
	return grid.Rows[y]
}

func (grid *Grid) GetColumn(x int) string {
	result := make([]rune, grid.ColumnLength())
	for y := 0; y < grid.ColumnLength(); y++ {
		result[y] = grid.GetCharAtXY(x, y)
	}
	return string(result)
}

func (grid *Grid) ColumnLength() int {
	return len(grid.Rows)
}

func (grid *Grid) RowLength() int {
	return len(grid.Rows[0])
}

func (grid *Grid) GetAdjacentPoints(point Point) []Point {
	result := make([]Point, 0)
	isValid := func(point Point) bool {
		return point.X >= 0 && point.Y >= 0 && point.X < grid.RowLength() && point.Y < grid.ColumnLength()
	}
	potentialResult := []Point{
		{point.X - 1, point.Y},
		{point.X - 1, point.Y - 1},
		{point.X - 1, point.Y + 1},
		{point.X, point.Y - 1},
		{point.X, point.Y + 1},
		{point.X + 1, point.Y},
		{point.X + 1, point.Y - 1},
		{point.X + 1, point.Y + 1},
	}
	for _, point := range potentialResult {
		if isValid(point) {
			result = append(result, point)
		}
	}
	return result
}

func (grid *Grid) GetAdjacents(point Point, diagonal bool) []Point {
	if diagonal {
		return grid.GetAdjacentPoints(point)
	} else {
		result := make([]Point, 0)
		isValid := func(point Point) bool {
			return point.X >= 0 && point.Y >= 0 && point.X < grid.RowLength() && point.Y < grid.ColumnLength()
		}
		potentialResult := []Point{
			{point.X - 1, point.Y},
			{point.X + 1, point.Y},
			{point.X, point.Y - 1},
			{point.X, point.Y + 1},
		}
		for _, point := range potentialResult {
			if isValid(point) {
				result = append(result, point)
			}
		}
		return result
	}
}

func (grid *Grid) InsertRow(rowNumber int, char rune) {

	newRow := make([]rune, len(grid.Rows[0]))
	for i := 0; i < len(grid.Rows[0]); i++ {
		newRow[i] = char
	}
	grid.Rows = append(grid.Rows[:rowNumber+1], grid.Rows[rowNumber:]...)
	grid.Rows[rowNumber] = string(newRow)
}

func (grid *Grid) InsertColumn(columnNumber int, char rune) {
	for rowNumber := 0; rowNumber < grid.ColumnLength(); rowNumber++ {
		grid.Rows[rowNumber] = grid.Rows[rowNumber][:columnNumber] + string(char) + grid.Rows[rowNumber][columnNumber:]
	}
}

func (p1 *Point) DistanceTo(p2 Point) int {
	return int(math.Abs(float64(p1.X)-float64(p2.X)) + math.Abs(float64(p1.Y)-float64(p2.Y)))
}

func (g *Grid) Draw() {
	for _, line := range g.Rows {
		fmt.Println(line)
	}
}

func (g *Grid) IsValid(point Point) bool {
	return point.X >= 0 && point.Y >= 0 && point.X < g.RowLength() && point.Y < g.ColumnLength()
}
