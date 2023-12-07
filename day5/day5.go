package day5

import (
	"fmt"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
	"github.com/schollz/progressbar/v3"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day5/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day5/input/input1.txt")
}

type AlMap struct {
	start   int
	end     int
	mapping int
}

func ReadAlMap(line string) AlMap {
	lineAsFields := strings.Fields(line)
	startInt, _ := strconv.Atoi(lineAsFields[1])
	sizeInt, _ := strconv.Atoi(lineAsFields[2])
	mappingInt, _ := strconv.Atoi(lineAsFields[0])
	return AlMap{
		start:   startInt,
		end:     startInt + sizeInt,
		mapping: mappingInt,
	}
}

func (m *AlMap) Apply(value int) (int, bool) {
	if value >= m.start && value <= m.end {
		return m.mapping + (value - m.start), true
	}
	return -1, false
}

func MapAlValue(value int, mappings []AlMap) int {
	result := value
	for _, mapping := range mappings {
		if mappedValue, ok := mapping.Apply(result); ok {
			return mappedValue
		}
	}
	return result
}

func SolvePart1(inputFile string) string {
	result := -1
	seeds := make([]int, 0)
	allMappings := make([][]AlMap, 0)
	currentMapping := make([]AlMap, 0)

	for _, line := range helper.ReadInputFile(inputFile) {
		if strings.HasPrefix(line, "seeds: ") {
			for _, seedString := range strings.Fields(strings.Replace(line, "seeds: ", "", 1)) {
				seedInt, _ := strconv.Atoi(seedString)
				seeds = append(seeds, seedInt)
			}
		} else if strings.Contains(line, "-to-") {
			allMappings = append(allMappings, currentMapping)
			currentMapping = make([]AlMap, 0)
		} else if line != "" {
			currentMapping = append(currentMapping, ReadAlMap(line))
		}
	}
	allMappings = append(allMappings, currentMapping)

	for _, seed := range seeds {
		mappedSeedValue := seed
		for _, mapping := range allMappings {
			mappedSeedValue = MapAlValue(mappedSeedValue, mapping)
		}
		if result == -1 || result > mappedSeedValue {
			result = mappedSeedValue
		}
	}

	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	result := -1
	seeds := make([]int, 0)
	allMappings := make([][]AlMap, 0)
	currentMapping := make([]AlMap, 0)

	for _, line := range helper.ReadInputFile(inputFile) {
		if strings.HasPrefix(line, "seeds: ") {
			for _, seedString := range strings.Fields(strings.Replace(line, "seeds: ", "", 1)) {
				seedInt, _ := strconv.Atoi(seedString)
				seeds = append(seeds, seedInt)
			}
		} else if strings.Contains(line, "-to-") {
			allMappings = append(allMappings, currentMapping)
			currentMapping = make([]AlMap, 0)
		} else if line != "" {
			currentMapping = append(currentMapping, ReadAlMap(line))
		}
	}
	allMappings = append(allMappings, currentMapping)

	totalSeedsToTry, countOfTriedSeeds := 0, 0
	for i := 1; i < len(seeds); i = i + 2 {
		totalSeedsToTry += seeds[i]
	}
	bar := progressbar.Default(int64(totalSeedsToTry))

	for i := 0; i < len(seeds); i = i + 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			mappedSeedValue := seed
			for _, mapping := range allMappings {
				mappedSeedValue = MapAlValue(mappedSeedValue, mapping)
			}
			if result == -1 || result > mappedSeedValue {
				result = mappedSeedValue
			}
			countOfTriedSeeds++
			if countOfTriedSeeds%10000 == 0 {
				bar.Add(10000)
			}
		}
	}
	bar.Finish()

	return fmt.Sprint(result)
}
