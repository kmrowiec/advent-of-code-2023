package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Solver interface {
	PartOne() string
	PartTwo() string
}

func ReadInputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([]string, 0)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

type Point struct {
	X, Y int
}

type Grid struct {
	Rows []string
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

func (grid *Grid) GerRow(y int) string {
	return grid.Rows[y]
}

func (grid *Grid) Lentgh() int {
	return len(grid.Rows)
}

func (grid *Grid) LentghX() int {
	return len(grid.Rows[0])
}

func (grid *Grid) GetAdjacentPoints(point Point) []Point {
	result := make([]Point, 0)
	isValid := func(point Point) bool {
		return point.X >= 0 && point.Y >= 0 && point.X < grid.LentghX() && point.Y < grid.Lentgh()
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
			return point.X >= 0 && point.Y >= 0 && point.X < grid.LentghX() && point.Y < grid.Lentgh()
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

func (g *Grid) Draw() {
	for _, line := range g.Rows {
		fmt.Println(line)
	}
}
