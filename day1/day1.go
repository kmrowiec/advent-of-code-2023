package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day1/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day1/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	var result int

	for _, line := range helper.ReadInputFile(inputFile) {
		numbersOnly := make([]rune, 0)
		for _, c := range line {
			if unicode.IsNumber(c) {
				numbersOnly = append(numbersOnly, c)
			}
		}
		textValue := "" + string(numbersOnly[0]) + string(numbersOnly[len(numbersOnly)-1])
		intValue, _ := strconv.Atoi(textValue)
		result += intValue
	}

	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	var result int

	for _, line := range helper.ReadInputFile(inputFile) {
		numbersOnly := FindOnlyNumbers(line)
		lineSumAsText := string(numbersOnly[0]) + string(numbersOnly[len(numbersOnly)-1])
		log.Println(line + " -> " + numbersOnly + " / " + lineSumAsText)
		lineSumAsInt, _ := strconv.Atoi(lineSumAsText)
		result += lineSumAsInt
	}

	return fmt.Sprint(result)
}

func FindOnlyNumbers(line string) string {
	type pair struct {
		text   string
		number string
	}
	spelledNumbers := []pair{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}

	result := ""
	for i, c := range line {
		for _, spelledNumber := range spelledNumbers {
			if strings.HasPrefix(string(line[i:]), spelledNumber.text) {
				result += spelledNumber.number
			}
		}
		if unicode.IsNumber(c) {
			result += string(c)
		}
	}
	return result
}
