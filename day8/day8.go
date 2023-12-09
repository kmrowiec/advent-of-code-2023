package day8

import (
	"fmt"
	"strings"

	"dev.kmrowiec/aoc/helper"
	"github.com/schollz/progressbar/v3"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day8/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day8/input/input1.txt")
}

type Node struct {
	Left  string
	Right string
}

func ReadNode(line string) (string, Node) {
	idAndNode := strings.Split(line, " = ")
	id := idAndNode[0]
	leftAndRightString := strings.Split(idAndNode[1], ", ")
	return id, Node{
		Left:  leftAndRightString[0][1:],
		Right: leftAndRightString[1][:len(leftAndRightString)+1],
	}
}

func SolvePart1(inputFile string) string {
	result := 0
	lines := helper.ReadInputFile(inputFile)
	instructions := lines[0]
	nodes := map[string]Node{}
	for _, line := range lines[2:] {
		id, node := ReadNode(line)
		nodes[id] = node
	}

	currentNode := "AAA"
	for currentInstruction := 0; ; currentInstruction++ {
		if currentNode == "ZZZ" {
			return fmt.Sprint(result)
		}
		if currentInstruction >= len(instructions) {
			currentInstruction = 0
		}
		inst := instructions[currentInstruction]
		if inst == 'L' {
			currentNode = nodes[currentNode].Left
		} else {
			currentNode = nodes[currentNode].Right
		}
		result++
	}
}

func SolvePart2(inputFile string) string {
	result := 0
	lines := helper.ReadInputFile(inputFile)
	instructions := lines[0]
	nodes := map[string]Node{}
	for _, line := range lines[2:] {
		id, node := ReadNode(line)
		nodes[id] = node
	}

	currentNodes := make([]string, 0)

	for key := range nodes {
		if key[2] == 'A' {
			currentNodes = append(currentNodes, key)
		}
	}

	bar := progressbar.Default(-1)

	completedMap := map[string]int{}
	for i := 0; ; i++ {

		if i >= len(instructions) {
			i = 0
		}

		inst := instructions[i]
		if inst == 'L' {
			for index, id := range currentNodes {
				currentNodes[index] = nodes[id].Left
			}
		} else {
			for index, id := range currentNodes {
				currentNodes[index] = nodes[id].Right
			}
		}

		result++
		// fmt.Println(currentNodes)

		for _, id := range currentNodes {
			if id[2] == 'Z' {
				if _, ok := completedMap[id]; !ok {
					completedMap[id] = result
				}
			}
		}
		if len(completedMap) == len(currentNodes) {
			break
		}
		bar.Add(1)
	}

	completedIndexes := make([]int, 0)
	for _, v := range completedMap {
		completedIndexes = append(completedIndexes, v)
	}

	fmt.Println(completedIndexes)

	// TODO find lowest common multiplier for the completed indexes

	return fmt.Sprint(result)
}
