package day18

import (
	"fmt"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
	"dev.kmrowiec/aoc/helper/grid"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day18/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day18/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	fields := make([][]string, 0)
	for _, line := range helper.ReadInputFile(inputFile) {
		fields = append(fields, strings.Fields(line))
	}
	return fmt.Sprint(calculateArea(fields))
}

func calculateArea(instructions [][]string) (result int) {
	result = 0

	boundaryPointCount := 0
	x, y := 0, 0
	points := make([]grid.Point, 0)

	for _, instruction := range instructions {
		direction := instruction[0]
		length, _ := strconv.Atoi(instruction[1])

		switch direction {
		case "R":
			x += length
		case "L":
			x -= length
		case "U":
			y -= length
		case "D":
			y += length
		}
		boundaryPointCount += length
		points = append(points, grid.PointFromXY(x, y))
	}

	points = append(points, points[0])

	sum := 0
	for i := 0; i < len(points)-1; i++ {
		p1, p2 := points[i], points[i+1]
		sum += p1.X*p2.Y - p1.Y*p2.X
	}

	area := sum / 2
	result = area + boundaryPointCount/2 + 1
	return
}

func SolvePart2(inputFile string) string {
	fields := make([][]string, 0)
	for _, line := range helper.ReadInputFile(inputFile) {
		hexValue := strings.TrimSuffix(strings.TrimPrefix(strings.Fields(line)[2], "("), ")")
		fields = append(fields, DecodeInstruction(hexValue))
	}
	return fmt.Sprint(calculateArea(fields))
}

func DecodeInstruction(hexInput string) []string {
	result := make([]string, 2)
	encodedDistance := hexInput[1:6]
	encodedDirection := hexInput[6:]

	switch encodedDirection {
	case "0":
		result[0] = "R"
	case "1":
		result[0] = "D"
	case "2":
		result[0] = "L"
	case "3":
		result[0] = "U"
	}

	decimalDistance, _ := strconv.ParseInt(encodedDistance, 16, 64)
	result[1] = fmt.Sprint(decimalDistance)
	return result
}
