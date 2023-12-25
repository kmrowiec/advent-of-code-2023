package day17

import (
	"fmt"
	"strconv"

	"dev.kmrowiec/aoc/helper/grid"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day17/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day17/input/input1.txt")
}

type Node struct {
	Point     grid.Point
	Direction rune
	Repeated  int
}

func (n *Node) String() string {
	return fmt.Sprintf("(%d,%d) %c%d", n.Point.X, n.Point.Y, n.Direction, n.Repeated)
}

func SolvePart1(inputFile string) string {
	result := 999999
	g := grid.GridFromFile(inputFile)

	scores := map[Node]int{}
	visitedNodes := map[Node]bool{}
	unvisitedSet := map[Node]bool{}

	currentNode := Node{Point: grid.PointFromXY(0, 0), Direction: 'X', Repeated: 0}
	scores[currentNode] = 0
	unvisitedSet[currentNode] = true

	neighbours := findUnvisitedNeighbours(currentNode, visitedNodes, &g)

	for len(unvisitedSet) > 0 {
		for _, node := range neighbours {

			// Calculate score for the neighbour
			fmt.Printf("Calculating score for node (%d,%d), %c,%d | ", node.Point.X, node.Point.Y, node.Direction, node.Repeated)
			newScore := scores[currentNode]
			max := node.Repeated
			if currentNode.Direction == node.Direction {
				max = node.Repeated - currentNode.Repeated
			}
			for i := 1; i <= max; i++ {
				switch node.Direction {
				case 'E':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X+i, currentNode.Point.Y), &g)
				case 'W':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X-i, currentNode.Point.Y), &g)
				case 'S':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y+i), &g)
				case 'N':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y-i), &g)
				}
			}
			fmt.Printf("Score is %d\n", newScore)

			// Save the new score
			if value, ok := scores[node]; ok {
				if newScore < value {
					scores[node] = newScore
				}
			} else {
				scores[node] = newScore
			}

			// If the node is the end, we remember the score if lowest so far
			if node.Point == grid.PointFromXY(g.RowLength()-1, g.ColumnLength()-1) {
				if result > newScore {
					result = newScore
				}
			}
		}

		visitedNodes[currentNode] = true
		delete(unvisitedSet, currentNode)

		// Add all neighbours to unvisited so we can iterate over them next
		for _, n := range neighbours {
			if _, ok := unvisitedSet[n]; !ok {
				unvisitedSet[n] = true
			}
		}

		var nextNode Node
		maxScore := 9999999

		for n := range unvisitedSet {
			nextNode = n
			break
		}

		for n := range unvisitedSet {
			if score, ok := scores[n]; ok {
				if score < maxScore {
					maxScore = score
					nextNode = n
				}
			}
		}

		neighbours = findUnvisitedNeighbours(nextNode, visitedNodes, &g)
		currentNode = nextNode

	}

	fmt.Printf("%-v", scores)

	return fmt.Sprint(result)
}

func getScore(point grid.Point, g *grid.Grid) (score int) {
	score, _ = strconv.Atoi(string(g.GetCharAt(point)))
	return
}

func findUnvisitedNeighbours(currentNode Node, visitedNodes map[Node]bool, g *grid.Grid) []Node {
	result := make([]Node, 0)

	if currentNode.Direction != 'W' {
		// East
		maxEast := 3
		eastRepeated := 0
		if currentNode.Direction == 'E' {
			maxEast = 3 - currentNode.Repeated
			eastRepeated = currentNode.Repeated
		}
		for i := 1; i <= maxEast; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X+i, currentNode.Point.Y), Direction: 'E', Repeated: eastRepeated + i})
		}
	}

	if currentNode.Direction != 'E' {
		// West
		max := 3
		repeated := 0
		if currentNode.Direction == 'W' {
			max = 3 - currentNode.Repeated
			repeated = currentNode.Repeated
		}
		for i := 1; i <= max; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X-i, currentNode.Point.Y), Direction: 'W', Repeated: repeated + i})
		}
	}

	if currentNode.Direction != 'N' {
		// South
		max := 3
		repeated := 0
		if currentNode.Direction == 'S' {
			max = 3 - currentNode.Repeated
			repeated = currentNode.Repeated
		}
		for i := 1; i <= max; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y+i), Direction: 'S', Repeated: repeated + i})
		}
	}

	if currentNode.Direction != 'S' {
		// North
		max := 3
		repeated := 0
		if currentNode.Direction == 'N' {
			max = 3 - currentNode.Repeated
			repeated = currentNode.Repeated
		}
		for i := 1; i <= max; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y-i), Direction: 'N', Repeated: repeated + i})
		}
	}

	filteredResults := make([]Node, 0)
	for _, node := range result {
		if g.IsValid(node.Point) {
			if _, ok := visitedNodes[node]; !ok {
				filteredResults = append(filteredResults, node)
			}
		}
	}

	fmt.Println("Current "+currentNode.String()+".Found unvisited neighbours: ", filteredResults)
	return filteredResults
}

func SolvePart2(inputFile string) string {
	result := 999999
	g := grid.GridFromFile(inputFile)

	scores := map[Node]int{}
	visitedNodes := map[Node]bool{}
	unvisitedSet := map[Node]bool{}

	currentNode := Node{Point: grid.PointFromXY(0, 0), Direction: 'E', Repeated: 0}
	nextNode := Node{Point: grid.PointFromXY(0, 0), Direction: 'S', Repeated: 0}
	scores[currentNode] = 0
	scores[nextNode] = 0
	unvisitedSet[currentNode] = true
	unvisitedSet[nextNode] = true

	neighbours := findUnvisitedNeighboursPartTwo(currentNode, visitedNodes, &g)

	for len(unvisitedSet) > 0 {

		for _, node := range neighbours {

			// Calculate score for the neighbour
			// fmt.Printf("Calculating score for node (%d,%d), %c,%d | ", node.Point.X, node.Point.Y, node.Direction, node.Repeated)
			newScore := scores[currentNode]
			max := node.Repeated
			if currentNode.Direction == node.Direction {
				max = node.Repeated - currentNode.Repeated
			}
			for i := 1; i <= max; i++ {
				switch node.Direction {
				case 'E':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X+i, currentNode.Point.Y), &g)
				case 'W':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X-i, currentNode.Point.Y), &g)
				case 'S':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y+i), &g)
				case 'N':
					newScore += getScore(grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y-i), &g)
				}
			}
			// fmt.Printf("Score is %d\n", newScore)

			// Save the new score
			if value, ok := scores[node]; ok {
				if newScore < value {
					scores[node] = newScore
				}
			} else {
				scores[node] = newScore
			}

			// If the node is the end, we remember the score if lowest so far
			if node.Point == grid.PointFromXY(g.RowLength()-1, g.ColumnLength()-1) {
				if result > newScore {
					result = newScore
				}
			}
		}

		visitedNodes[currentNode] = true
		delete(unvisitedSet, currentNode)

		// Add all neighbours to unvisited so we can iterate over them next
		for _, n := range neighbours {
			if _, ok := unvisitedSet[n]; !ok {
				unvisitedSet[n] = true
			}
		}

		var nextNode Node
		maxScore := 9999999

		for n := range unvisitedSet {
			nextNode = n
			break
		}

		for n := range unvisitedSet {
			if score, ok := scores[n]; ok {
				if score < maxScore {
					maxScore = score
					nextNode = n
				}
			}
		}

		neighbours = findUnvisitedNeighboursPartTwo(nextNode, visitedNodes, &g)
		currentNode = nextNode

	}

	fmt.Printf("%-v", scores)

	return fmt.Sprint(result)
}

func findUnvisitedNeighboursPartTwo(currentNode Node, visitedNodes map[Node]bool, g *grid.Grid) []Node {
	result := make([]Node, 0)

	if currentNode.Direction == 'E' || ((currentNode.Direction == 'S' || currentNode.Direction == 'N') && currentNode.Repeated > 3) {
		// East
		maxEast := 10
		eastRepeated := 0
		if currentNode.Direction == 'E' {
			maxEast = 10 - currentNode.Repeated
			eastRepeated = currentNode.Repeated
		}
		for i := 4; i <= maxEast; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X+i, currentNode.Point.Y), Direction: 'E', Repeated: eastRepeated + i})
		}
	}

	if currentNode.Direction == 'W' || ((currentNode.Direction == 'S' || currentNode.Direction == 'N') && currentNode.Repeated > 3) {
		// West
		max := 10
		repeated := 0
		if currentNode.Direction == 'W' {
			max = 10 - currentNode.Repeated
			repeated = currentNode.Repeated
		}
		for i := 4; i <= max; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X-i, currentNode.Point.Y), Direction: 'W', Repeated: repeated + i})
		}
	}

	if currentNode.Direction == 'S' || ((currentNode.Direction == 'E' || currentNode.Direction == 'W') && currentNode.Repeated > 3) {
		// South
		max := 10
		repeated := 0
		if currentNode.Direction == 'S' {
			max = 10 - currentNode.Repeated
			repeated = currentNode.Repeated
		}
		for i := 4; i <= max; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y+i), Direction: 'S', Repeated: repeated + i})
		}
	}

	if currentNode.Direction == 'N' || ((currentNode.Direction == 'E' || currentNode.Direction == 'W') && currentNode.Repeated > 3) {
		// North
		max := 10
		repeated := 0
		if currentNode.Direction == 'N' {
			max = 10 - currentNode.Repeated
			repeated = currentNode.Repeated
		}
		for i := 4; i <= max; i++ {
			result = append(result, Node{Point: grid.PointFromXY(currentNode.Point.X, currentNode.Point.Y-i), Direction: 'N', Repeated: repeated + i})
		}
	}

	filteredResults := make([]Node, 0)
	for _, node := range result {
		if g.IsValid(node.Point) {
			if _, ok := visitedNodes[node]; !ok {
				filteredResults = append(filteredResults, node)
			}
		}
	}

	fmt.Println("Current "+currentNode.String()+". Found unvisited neighbours: ", filteredResults)
	return filteredResults
}
