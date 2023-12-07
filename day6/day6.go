package day6

import (
	"fmt"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day6/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day6/input/input1.txt")
}

type RaceRecord struct {
	Time     int
	Distance int
}

func (r *RaceRecord) GetMarginOfError() int {
	lowerBoundary, upperBoundary := -1, -1
	for i := 0; i < r.Time; i++ {
		scoreAtI := r.scoreIfReleasedAt(i)
		if scoreAtI > r.Distance {
			if lowerBoundary == -1 {
				lowerBoundary = i
			}
			upperBoundary = i
		}
	}
	return upperBoundary - lowerBoundary + 1
}

func (raceRecord *RaceRecord) scoreIfReleasedAt(time int) int {
	const acceleration = 1 // mm/ms^2
	speedAt := func(time int) int {
		return time * acceleration
	}
	remainingTime := raceRecord.Time - time
	return speedAt(time) * remainingTime
}

func LoadRaceRecords(lines []string) []RaceRecord {
	result := make([]RaceRecord, 0)
	timesAsStrings := strings.Fields(lines[0])[1:]
	distancesAsStrings := strings.Fields(lines[1])[1:]
	for i := 0; i < len(timesAsStrings); i++ {
		time, _ := strconv.Atoi(timesAsStrings[i])
		distance, _ := strconv.Atoi(distancesAsStrings[i])
		result = append(result, RaceRecord{Time: time, Distance: distance})
	}
	return result
}

func LoadRaceRecordsPartTwo(lines []string) RaceRecord {
	timesAsStrings := strings.Fields(lines[0])[1:]
	distancesAsStrings := strings.Fields(lines[1])[1:]
	time, _ := strconv.Atoi(strings.Join(timesAsStrings, ""))
	distance, _ := strconv.Atoi(strings.Join(distancesAsStrings, ""))
	return RaceRecord{Time: time, Distance: distance}
}

func SolvePart1(inputFile string) string {
	result := 1
	raceRecords := LoadRaceRecords(helper.ReadInputFile(inputFile))
	for _, raceRecord := range raceRecords {
		result *= raceRecord.GetMarginOfError()
	}
	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	result := 0
	raceRecord := LoadRaceRecordsPartTwo(helper.ReadInputFile(inputFile))
	result = raceRecord.GetMarginOfError()
	return fmt.Sprint(result)
}
