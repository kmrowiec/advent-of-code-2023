package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
	"golang.org/x/exp/maps"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day7/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day7/input/input1.txt")
}

type Hand struct {
	Raw      string
	Strength int
	Bet      int
}

func ReadHand(line string) Hand {
	splitLine := strings.Fields(line)
	bet, _ := strconv.Atoi(splitLine[1])
	return Hand{
		Raw:      splitLine[0],
		Strength: CalculateStrength(splitLine[0]),
		Bet:      bet}
}

func ReadHandWithJokers(line string) Hand {
	splitLine := strings.Fields(line)
	bet, _ := strconv.Atoi(splitLine[1])
	return Hand{
		Raw:      splitLine[0],
		Strength: CalculateStrengthWithJokers(splitLine[0]),
		Bet:      bet}
}

func CalculateStrength(rawHand string) int {
	cardMap := map[rune]int{}
	for _, c := range rawHand {
		if val, ok := cardMap[c]; ok {
			cardMap[c] = val + 1
		} else {
			cardMap[c] = 1
		}
	}
	switch len(cardMap) {
	case 5:
		return 6 // high pair
	case 4:
		return 5 // pair
	case 3:
		values := maps.Values(cardMap)
		if values[0] == 3 || values[1] == 3 || values[2] == 3 {
			return 3 // three of a kind
		} else {
			return 4 // two pair
		}
	case 2:
		// if split is 3/2 then full house
		// if split is 4/1 then foak
		values := maps.Values(cardMap)
		if values[0] == 1 || values[1] == 1 {
			return 1 // four of a kind
		} else {
			return 2 // full house
		}
	case 1:
		return 0 // five of a kind
	}
	return 0
}

func CalculateStrengthWithJokers(rawHand string) int {
	prevStrength := CalculateStrength(rawHand)
	jokerCount := strings.Count(rawHand, "J")

	switch jokerCount {
	case 1:
		if prevStrength == 5 || prevStrength == 4 || prevStrength == 3 {
			return prevStrength - 2
		} else {
			return prevStrength - 1
		}
	case 2:
		if prevStrength == 4 {
			return 1
		} else {
			return prevStrength - 2
		}
	case 3:
		return prevStrength - 2
	case 4:
		return 0
	}

	return prevStrength
}

type ByRaw []Hand

func (a ByRaw) Len() int           { return len(a) }
func (a ByRaw) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRaw) Less(i, j int) bool { return CompareRawHands(a[i], a[j]) }

type ByRawWithJokers []Hand

func (a ByRawWithJokers) Len() int           { return len(a) }
func (a ByRawWithJokers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRawWithJokers) Less(i, j int) bool { return CompareRawHandsWithJokers(a[i], a[j]) }

var cardOrder = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

var cardOrderWithJokers = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

func CompareRawHands(first, second Hand) bool {
	for i := range first.Raw {
		a, b := cardOrder[rune(first.Raw[i])], cardOrder[rune(second.Raw[i])]
		if a != b {
			if a > b {
				return false
			} else {
				return true
			}
		}
	}
	return true
}

func CompareRawHandsWithJokers(first, second Hand) bool {
	for i := range first.Raw {
		a, b := cardOrderWithJokers[rune(first.Raw[i])], cardOrderWithJokers[rune(second.Raw[i])]
		if a != b {
			if a > b {
				return false
			} else {
				return true
			}
		}
	}
	return true
}

func SolvePart1(inputFile string) string {
	result := 0

	strengthsMap := map[int][]Hand{}
	for _, line := range helper.ReadInputFile(inputFile) {
		hand := ReadHand(line)
		if val, ok := strengthsMap[hand.Strength]; ok {
			strengthsMap[hand.Strength] = append(val, hand)
		} else {
			strengthsMap[hand.Strength] = []Hand{hand}
		}
	}

	allHands := make([]Hand, 0)
	for i := 6; i >= 0; i-- {
		if hands, ok := strengthsMap[i]; ok {
			sort.Sort(ByRaw(hands))
			allHands = append(allHands, hands...)
		}
	}

	for i, hand := range allHands {
		result += (i + 1) * hand.Bet
	}

	return fmt.Sprint(result)
}

func SolvePart2(inputFile string) string {
	result := 0

	strengthsMap := map[int][]Hand{}
	for _, line := range helper.ReadInputFile(inputFile) {
		hand := ReadHandWithJokers(line)
		if val, ok := strengthsMap[hand.Strength]; ok {
			strengthsMap[hand.Strength] = append(val, hand)
		} else {
			strengthsMap[hand.Strength] = []Hand{hand}
		}
	}

	allHands := make([]Hand, 0)
	for i := 6; i >= 0; i-- {
		if hands, ok := strengthsMap[i]; ok {
			sort.Sort(ByRawWithJokers(hands))
			allHands = append(allHands, hands...)
		}
	}

	fmt.Println(allHands)

	for i, hand := range allHands {
		result += (i + 1) * hand.Bet
	}

	return fmt.Sprint(result)
}
