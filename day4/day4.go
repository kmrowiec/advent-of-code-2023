package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day4/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day4/input/input1.txt")
}

type Scratchcard struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
	Copies         int
}

func (card *Scratchcard) CountWinningNumbers() int {
	winningMap := map[int]bool{}
	result := 0
	for _, number := range card.WinningNumbers {
		winningMap[number] = true
	}
	for _, number := range card.Numbers {
		if _, ok := winningMap[number]; ok {
			result++
		}
	}
	return result
}

func ReadScratchcard(line string) Scratchcard {
	result := Scratchcard{}
	idAndRest := strings.Split(line, ": ")
	idAsString := strings.Replace(idAndRest[0], "Card ", "", 1)
	result.Id, _ = strconv.Atoi(idAsString)

	numbersPart := strings.Split(idAndRest[1], " | ")
	winningPart := strings.Fields(numbersPart[0])
	actualPart := strings.Fields(numbersPart[1])

	for _, numberAsString := range winningPart {
		number, _ := strconv.Atoi(numberAsString)
		result.WinningNumbers = append(result.WinningNumbers, number)
	}

	for _, numberAsString := range actualPart {
		number, _ := strconv.Atoi(numberAsString)
		result.Numbers = append(result.Numbers, number)
	}

	result.Copies = 1

	return result
}

func SolvePart1(inputFile string) string {
	scratchcards := make([]Scratchcard, 0)
	for _, line := range helper.ReadInputFile(inputFile) {
		scratchcards = append(scratchcards, ReadScratchcard(line))
	}
	result := 0
	for _, card := range scratchcards {
		winningCount := card.CountWinningNumbers()
		if winningCount > 0 {
			score := int(math.Pow(2, float64(winningCount)-1))
			result += score
		}
	}
	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	scratchcards := make([]Scratchcard, 0)
	for _, line := range helper.ReadInputFile(inputFile) {
		scratchcards = append(scratchcards, ReadScratchcard(line))
	}
	result := 0
	for index, card := range scratchcards {
		winningCount := card.CountWinningNumbers()
		if winningCount > 0 {
			for i := 0; i < winningCount; i++ {
				scratchcards[index+i+1].Copies += card.Copies
			}
		}
		result += card.Copies
	}
	return fmt.Sprint(result)
}
