package helper

import (
	"bufio"
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

func (grid *Grid) GerRow(y int) string {
	return grid.Rows[y]
}

func (grid *Grid) Lentgh() int {
	return len(grid.Rows)
}

func (grid *Grid) GetAdjacentPoints(point Point) []Point {
	result := make([]Point, 0)
	isValid := func(point Point) bool {
		return point.X >= 0 && point.Y >= 0 && point.X < grid.Lentgh() && point.Y < grid.Lentgh()
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
