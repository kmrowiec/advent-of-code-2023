package day9

import (
	"fmt"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day9/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day9/input/input1.txt")
}

func PredictNext(sequence []int) int {
	if len(sequence) == 1 {
		return sequence[0]
	}
	gaps := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		gaps[i] = sequence[i+1] - sequence[i]
	}
	gapMap := map[int]bool{}
	for _, v := range gaps {
		gapMap[v] = true
	}
	if len(gapMap) == 1 {
		return sequence[len(sequence)-1] + gaps[0]
	} else {
		return sequence[len(sequence)-1] + PredictNext(gaps)
	}
}

func readSequence(line string, revert bool) []int {
	sequence := make([]int, 0)

	for _, v := range strings.Fields(line) {
		number, _ := strconv.Atoi(v)
		sequence = append(sequence, number)
	}

	if !revert {
		return sequence
	} else {
		reverted := make([]int, len(sequence))
		for i := len(sequence) - 1; i >= 0; i-- {
			reverted[len(sequence)-1-i] = sequence[i]
		}
		return reverted
	}
}

func SolvePart1(inputFile string) string {
	result := 0
	for _, line := range helper.ReadInputFile(inputFile) {
		sequence := readSequence(line, false)
		fmt.Println(sequence)
		result += PredictNext(sequence)
	}
	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	result := 0
	for _, line := range helper.ReadInputFile(inputFile) {
		sequence := readSequence(line, true)
		fmt.Println(sequence)
		result += PredictNext(sequence)
	}
	return fmt.Sprint(result)
}
