package day12

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day12/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day12/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	result := 0
	for _, line := range helper.ReadInputFile(inputFile) {
		splits := strings.Fields(line)
		result += CountOptions(splits[0], splits[1])
	}
	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	result := 0
	for _, line := range helper.ReadInputFile(inputFile) {
		splits := strings.Fields(line)
		result += CountOptionsUnfolded(splits[0], splits[1])
		fmt.Printf("Line %v done. Result: %d\n", line, result)
	}
	return fmt.Sprint(result)
}

func CountOptions(pattern, target string) (result int) {
	result = len(GenerateOptions(pattern, target))
	fmt.Printf("No of options for %v / %v is %d \n", pattern, target, result)
	return
}

func CountOptionsUnfolded(pattern, target string) (result int) {
	return CountOptions(Unfold(pattern, target))
}

func GenerateOptions(pattern, target string) []string {
	// if no unknowns, check if matches and return a one element slice
	// else
	//	check if pattern plausible
	//		if not. return empty list
	//		if yes, return a sum of slices from generate options calls to both next possible patterns

	targetSlice := make([]int, 0)
	for _, numberString := range strings.Split(target, ",") {
		numberInt, _ := strconv.Atoi(numberString)
		targetSlice = append(targetSlice, numberInt)
	}

	return GenerateOptionsInt(pattern, targetSlice)
}

func GenerateOptionsInt(pattern string, target []int) []string {
	if !strings.Contains(pattern, "?") {
		if IsMatchingInt(pattern, target) {
			return []string{pattern}
		}
		return []string{}
	} else {
		newPattern1 := strings.Replace(pattern, "?", ".", 1)
		newPattern2 := strings.Replace(pattern, "?", "#", 1)
		return append(GenerateOptionsInt(newPattern1, target), GenerateOptionsInt(newPattern2, target)...)
	}
}

func IsMatching(pattern, target string) bool {
	return PatternToTarget(pattern) == target
}

func IsMatchingInt(pattern string, target []int) bool {
	return reflect.DeepEqual(PatternToTargetInt(pattern), target)
}

func PatternToTarget(pattern string) string {
	result := ""
	pattern = strings.TrimLeft(pattern, ".")
	pattern = strings.TrimRight(pattern, ".")

	if strings.Contains(pattern, "?") {
		panic("Pattern " + pattern + " contains unknowns and cannot be verified")
	}

	buffer := rune(0)
	count := 0
	for _, char := range pattern + "." {
		if char == '#' {
			if buffer != '.' {
				count++
			} else {
				if count != 0 {
					result = result + ","
				}
				count = 1
			}
			buffer = '#'
		}
		if char == '.' {
			if buffer != '#' {
				count++
			} else {
				if count != 0 {
					result = result + fmt.Sprint(count)
				}
				count = 1
			}
			buffer = '.'
		}
	}
	return result
}

func PatternToTargetInt(pattern string) []int {
	targetSlice := make([]int, 0)
	buffer := rune(0)
	count := 0
	for _, char := range pattern + "." {
		if char == '#' {
			if buffer != '.' {
				count++
			} else {
				count = 1
			}
			buffer = '#'
		}
		if char == '.' {
			if buffer != '#' {
				count++
			} else {
				if count != 0 {
					targetSlice = append(targetSlice, count)
				}
				count = 1
			}
			buffer = '.'
		}
	}
	return targetSlice
}

func Unfold(pattern, target string) (newPattern, newTarget string) {
	const copies = 5

	patterns, targets := make([]string, 0), make([]string, 0)

	for i := 0; i < copies; i++ {
		patterns = append(patterns, pattern)
		targets = append(targets, target)
	}

	newPattern = strings.Join(patterns, "?")
	newTarget = strings.Join(targets, ",")
	return
}
