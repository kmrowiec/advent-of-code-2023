package day2

import (
	"fmt"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day2/input/input1.txt",
		map[string]int{"green": 13, "red": 12, "blue": 14})
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day2/input/input1.txt")
}

type Game struct {
	id            int
	maxCubeCounts map[string]int
}

func SolvePart1(inputFile string, criteria map[string]int) string {
	lines := helper.ReadInputFile(inputFile)
	sum := 0
	for _, line := range lines {
		game := ReadGame(line)
		if IsPossible(game, criteria) {
			sum += game.id
		}
	}
	return fmt.Sprint(sum)
}

func SolvePart2(inputFile string) string {
	lines := helper.ReadInputFile(inputFile)
	sum := 0
	for _, line := range lines {
		game := ReadGame(line)
		power := 1
		for _, count := range game.maxCubeCounts {
			power *= count
		}
		sum += power
	}
	return fmt.Sprint(sum)
}

func ReadGame(line string) Game {

	split1 := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.TrimSpace(strings.Split(split1[0], "Game ")[1]))

	rounds := strings.Split(split1[1], "; ")

	gameMap := map[string]int{}

	for _, round := range rounds {
		roundParts := strings.Split(round, ", ")
		for _, part := range roundParts {
			p := strings.Split(part, " ")
			count, _ := strconv.Atoi(p[0])
			colour := p[1]

			if value, exists := gameMap[colour]; exists {
				if value < count {
					gameMap[colour] = count
				}
			} else {
				gameMap[colour] = count
			}
		}
	}

	return Game{id, gameMap}
}

func IsPossible(game Game, bag map[string]int) bool {
	for colour, bagAmount := range bag {
		if gameCount, exists := game.maxCubeCounts[colour]; exists && gameCount > bagAmount {
			return false
		}
	}
	return true
}
