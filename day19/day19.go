package day19

import (
	"fmt"
	"strconv"
	"strings"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day19/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day19/input/input1.txt")
}

type Rule struct {
	Condition [3]int
	Output    string
}

type Part map[rune]int

func (w *Rule) Matches(part Part) bool {
	if w.Condition[0] == 0 {
		return true // no condition, so always true
	}

	field := rune(w.Condition[0])
	greater := rune(w.Condition[1]) == '>'
	value := w.Condition[2]

	if greater {
		return part[field] > value
	} else {
		return part[field] < value
	}
}

func ReadWorkflow(line string) (string, []Rule) {
	parts := strings.Split(line, "{")
	ruleString := strings.TrimSuffix(parts[1], "}")
	ruleParts := strings.Split(ruleString, ",")
	rules := make([]Rule, 0)
	for _, rulePart := range ruleParts {
		if conditionString, output, found := strings.Cut(rulePart, ":"); found {
			fieldName := int(conditionString[0])
			greaterOrNot := int(conditionString[1])
			value, _ := strconv.Atoi(conditionString[2:])
			rules = append(rules, Rule{Condition: [3]int{fieldName, greaterOrNot, value}, Output: output})
		} else {
			rules = append(rules, Rule{Condition: [3]int{}, Output: rulePart})
		}
	}
	return parts[0], rules
}

func ReadPart(line string) Part {
	var result Part = map[rune]int{}
	splitPart := strings.Split(line[1:len(line)-1], ",")
	for _, value := range splitPart {
		fieldName, fieldValue, _ := strings.Cut(value, "=")
		intFieldValue, _ := strconv.Atoi(fieldValue)
		result[rune(fieldName[0])] = intFieldValue
	}
	return result
}

func isDone(output string) bool {
	return output == "A" || output == "R"
}

func SolvePart1(inputFile string) string {
	result := 0

	lines := helper.ReadInputFile(inputFile)

	workflows := map[string][]Rule{}
	parts := make([]Part, 0)
	isWorkflow := true
	for _, line := range lines {
		if line == "" {
			isWorkflow = false
			continue
		}
		if isWorkflow {
			id, rules := ReadWorkflow(line)
			workflows[id] = rules
		} else {
			parts = append(parts, ReadPart(line))
		}
	}

	acceptedParts := make([]Part, 0)

	for _, part := range parts {
		output := "in"
		for !isDone(output) {
			rules := workflows[output]
			for _, rule := range rules {
				if rule.Matches(part) {
					output = rule.Output
					break
				}
			}
		}
		if output == "A" {
			acceptedParts = append(acceptedParts, part)
		}
	}

	fmt.Printf("Loaded %d workflows: %v\n", len(workflows), workflows)
	fmt.Printf("Loaded %d parts: %v\n", len(parts), parts)
	fmt.Printf("Found %d accepted parts: %v\n", len(acceptedParts), acceptedParts)

	for _, part := range acceptedParts {
		result += part['x'] + part['m'] + part['a'] + part['s']
	}

	return fmt.Sprint(result)
}

type Range struct {
	From, To int
}

func (r Range) GetLength() int {
	return r.To - r.From + 1
}

type RangeMap map[rune]Range

type AppliedRangeMap struct {
	workflowId string
	rangeMap   RangeMap
}

func (rangeMap RangeMap) ApplyTo(workflowId string, workflows map[string][]Rule) []AppliedRangeMap {
	rules := workflows[workflowId]
	result := make([]AppliedRangeMap, 0)

	fmt.Println("To workflow "+workflowId+", applying ranges ", rangeMap)

	currentRanges := rangeMap

	for _, rule := range rules {
		if rule.Condition[0] == 0 {
			result = append(result, AppliedRangeMap{workflowId: rule.Output, rangeMap: currentRanges})
		} else {
			matching, notMatching := splitRange(currentRanges, rule.Condition)
			if len(matching) > 0 {
				result = append(result, AppliedRangeMap{workflowId: rule.Output, rangeMap: matching})
			}
			if len(notMatching) == 0 {
				break
			}
			currentRanges = notMatching
		}
	}
	fmt.Println("Results: ", result)
	return result
}

func splitRange(rangeMap RangeMap, condition [3]int) (matching, notMatching RangeMap) {
	matching, notMatching = RangeMap{}, RangeMap{}

	for k, v := range rangeMap {
		matching[k] = v
		notMatching[k] = v
	}

	field := rune(condition[0])
	greater := rune(condition[1]) == '>'
	value := condition[2]

	r := rangeMap[field]

	if greater {
		if r.From > value {
			// Full match
			return rangeMap, RangeMap{}
		} else if r.To <= value {
			// Full mismatch
			return RangeMap{}, rangeMap
		} else {
			// split
			nonMatchingRange := Range{From: r.From, To: value}
			matchingRange := Range{From: value + 1, To: r.To}
			matching[field] = matchingRange
			notMatching[field] = nonMatchingRange
			return
		}
	} else {
		if r.To < value {
			// Full match
			return rangeMap, RangeMap{}
		} else if r.From >= value {
			// Full mismatch
			return RangeMap{}, rangeMap
		} else {
			// split
			matchingRange := Range{From: r.From, To: value - 1}
			nonMatchingRange := Range{From: value, To: r.To}
			matching[field] = matchingRange
			notMatching[field] = nonMatchingRange
			return
		}
	}
}

func SolvePart2(inputFile string) string {
	result := 0
	lines := helper.ReadInputFile(inputFile)

	workflows := map[string][]Rule{}
	for _, line := range lines {
		if line == "" {
			break
		}
		id, rules := ReadWorkflow(line)
		workflows[id] = rules
	}

	startingRanges := RangeMap{}
	startingRanges['x'] = Range{From: 1, To: 4000}
	startingRanges['m'] = Range{From: 1, To: 4000}
	startingRanges['a'] = Range{From: 1, To: 4000}
	startingRanges['s'] = Range{From: 1, To: 4000}

	queue := []AppliedRangeMap{{workflowId: "in", rangeMap: startingRanges}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, value := range current.rangeMap.ApplyTo(current.workflowId, workflows) {
			if value.workflowId == "A" {
				result += value.rangeMap['x'].GetLength() * value.rangeMap['m'].GetLength() * value.rangeMap['a'].GetLength() * value.rangeMap['s'].GetLength()
			} else if value.workflowId != "R" {
				queue = append(queue, value)
			}
		}
	}

	return fmt.Sprint(result)
}
