package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day3/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day3/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	grid := LoadGrid(inputFile)
	sum := 0
	for y := 0; y < grid.Lentgh(); y++ {
		for _, value := range FindNumbersInLine(grid, y) {
			sum += value
		}
	}
	return fmt.Sprint(sum)
}

func LoadGrid(inputFile string) helper.Grid {
	lines := helper.ReadInputFile(inputFile)
	grid := helper.Grid{Rows: lines}
	return grid
}

func FindNumbersInLine(grid helper.Grid, y int) []int {
	row := grid.GerRow(y)
	result := make([]int, 0)
	currentNumber := ""
	currentNumberIsValid := false
	for x, char := range row {
		if unicode.IsNumber(char) {
			currentNumber += string(char)
			currentNumberIsValid = currentNumberIsValid || validatePoisiton(helper.Point{X: x, Y: y}, grid)
		} else if currentNumber != "" {
			if currentNumberIsValid {
				n, _ := strconv.Atoi(currentNumber)
				result = append(result, n)
			}
			currentNumber = ""
			currentNumberIsValid = false
		}
	}
	if currentNumber != "" && currentNumberIsValid {
		n, _ := strconv.Atoi(currentNumber)
		result = append(result, n)
	}
	return result
}

func validatePoisiton(point helper.Point, grid helper.Grid) bool {
	for _, adjPoint := range grid.GetAdjacentPoints(point) {
		if p := grid.GetCharAt(adjPoint); p != '.' && !unicode.IsNumber(p) {
			return true
		}
	}
	return false
}

func SolvePart2(inputFile string) string {
	grid := LoadGrid(inputFile)
	cogMap := findCogsWithNumbers(grid)

	sum := 0
	for key, value := range cogMap {
		if len(value) == 2 {
			sum += value[0] * value[1]
		} else {
			fmt.Println(key, value)
		}
	}

	// fmt.Println(cogMap)
	return fmt.Sprint(sum)
}

func findCogsWithNumbers(grid helper.Grid) map[helper.Point][]int {
	result := map[helper.Point][]int{}
	cogs := map[helper.Point]bool{}
	for y := 0; y < grid.Lentgh(); y++ {
		currentNumber := ""
		for x, char := range grid.Rows[y] {
			if unicode.IsNumber(char) {
				currentNumber += string(char)
				for _, val := range findNearbyCogs(helper.Point{X: x, Y: y}, grid) {
					cogs[val] = true
				}
			} else if currentNumber != "" {
				n, _ := strconv.Atoi(currentNumber)
				for cog, _ := range cogs {
					if value, ok := result[cog]; ok {
						result[cog] = append(value, n)
					} else {
						result[cog] = []int{n}
					}
				}
				currentNumber = ""
				cogs = map[helper.Point]bool{}
			}
		}
		if currentNumber != "" {
			n, _ := strconv.Atoi(currentNumber)
			for cog, _ := range cogs {
				if value, ok := result[cog]; ok {
					result[cog] = append(value, n)
				} else {
					result[cog] = []int{n}
				}
			}
		}
		currentNumber = ""
		cogs = map[helper.Point]bool{}
	}
	return result
}

func findNearbyCogs(point helper.Point, grid helper.Grid) []helper.Point {
	cogs := make([]helper.Point, 0)
	for _, adjPoint := range grid.GetAdjacentPoints(point) {
		if p := grid.GetCharAt(adjPoint); p == '*' {
			cogs = append(cogs, adjPoint)
		}
	}
	return cogs
}
