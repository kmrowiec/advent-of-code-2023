package day15

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day15/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day15/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	result := 0
	inputLine := helper.ReadInputFile(inputFile)[0]
	steps := strings.Split(inputLine, ",")
	for _, step := range steps {
		result += ApplyHash(step)
	}
	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {

	steps := strings.Split(helper.ReadInputFile(inputFile)[0], ",")
	boxes := map[int][]Lens{}

	for _, step := range steps {
		if strings.HasSuffix(step, "-") {
			// Remove lense from a box
			label := step[:len(step)-1]
			boxNumber := ApplyHash(label)
			if lensesInBox, ok := boxes[boxNumber]; ok {
				removedIndex := -1
				for i, lens := range lensesInBox {
					if lens.Label == label {
						removedIndex = i
						break
					}
				}
				if removedIndex != -1 {
					if removedIndex == 0 && len(lensesInBox) == 1 {
						delete(boxes, boxNumber)
					} else {
						boxes[boxNumber] = append(boxes[boxNumber][:removedIndex], boxes[boxNumber][removedIndex+1:]...)
					}
				}
			}
		} else {
			// Replace or add lense to box
			label, focalString, _ := strings.Cut(step, "=")
			focalStrength, _ := strconv.Atoi(focalString)
			newLens := MakeLens(label, focalStrength)
			boxNumber := ApplyHash(label)

			if lensesInBox, ok := boxes[boxNumber]; ok {
				foundLens := -1
				for i, lens := range lensesInBox {
					if lens.Label == label {
						foundLens = i
						break
					}
				}
				if foundLens > -1 {
					boxes[boxNumber][foundLens] = newLens
				} else {
					boxes[boxNumber] = append(boxes[boxNumber], newLens)
				}
			} else {
				boxes[boxNumber] = []Lens{newLens}
			}
		}
		log.Println("After " + step)
		log.Println(boxes)
	}

	return fmt.Sprint(CalculateScore(boxes))
}

func ApplyHash(text string) int {
	result := 0
	for _, char := range text {
		result += int(char)
		result *= 17
		result = result % 256
	}
	return result
}

type Lens struct {
	Label       string
	FocalLength int
}

func MakeLens(label string, focalLength int) Lens {
	return Lens{FocalLength: focalLength, Label: label}
}

func CalculateScore(boxes map[int][]Lens) int {
	result := 0
	for i := 0; i < 256; i++ {
		if lenses, ok := boxes[i]; ok {
			for j, lens := range lenses {
				result += (1 + i) * (j + 1) * lens.FocalLength
			}
		}
	}
	return result
}
