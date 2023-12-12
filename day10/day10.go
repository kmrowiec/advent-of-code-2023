package day10

import (
	"fmt"

	"dev.kmrowiec/aoc/helper"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day10/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day10/input/input1.txt")
}

type Pipeline struct {
	Grid helper.Grid
}

func (p Pipeline) AreConnected(p1, p2 helper.Point) bool {
	pipe1, pipe2 := p.Grid.GetCharAt(p1), p.Grid.GetCharAt(p2)
	// fmt.Printf("Checking connectivity: %c and %c \n", pipe1, pipe2)
	switch pipe1 {

	case 'S':
		if p2.X == p1.X {
			if p2.Y == p1.Y-1 {
				return pipe2 == '|' || pipe2 == 'F' || pipe2 == '7' || pipe2 == 'S'
			}
			if p2.Y == p1.Y+1 {
				return pipe2 == '|' || pipe2 == 'L' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
		if p2.Y == p1.Y {
			if p2.X == p1.X-1 {
				return pipe2 == '-' || pipe2 == 'F' || pipe2 == 'L' || pipe2 == 'S'
			}
			if p2.X == p1.X+1 {
				return pipe2 == '-' || pipe2 == '7' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
	case '|':
		if p2.X == p1.X {
			if p2.Y == p1.Y-1 {
				return pipe2 == '|' || pipe2 == 'F' || pipe2 == '7' || pipe2 == 'S'
			}
			if p2.Y == p1.Y+1 {
				return pipe2 == '|' || pipe2 == 'L' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
	case '-':
		if p2.Y == p1.Y {
			if p2.X == p1.X-1 {
				return pipe2 == '-' || pipe2 == 'F' || pipe2 == 'L' || pipe2 == 'S'
			}
			if p2.X == p1.X+1 {
				return pipe2 == '-' || pipe2 == '7' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
	case 'L':
		if p2.Y == p1.Y {
			if p2.X == p1.X+1 {
				return pipe2 == '-' || pipe2 == '7' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
		if p2.Y == p1.Y-1 {
			if p2.X == p1.X {
				return pipe2 == '|' || pipe2 == 'F' || pipe2 == '7' || pipe2 == 'S'
			}
		}
	case 'J':
		if p2.Y == p1.Y {
			if p2.X == p1.X-1 {
				return pipe2 == '-' || pipe2 == 'F' || pipe2 == 'L' || pipe2 == 'S'
			}
		}
		if p2.Y == p1.Y-1 {
			if p2.X == p1.X {
				return pipe2 == '|' || pipe2 == 'F' || pipe2 == '7' || pipe2 == 'S'
			}
		}
	case '7':
		if p2.Y == p1.Y {
			if p2.X == p1.X-1 {
				return pipe2 == '-' || pipe2 == 'F' || pipe2 == 'L' || pipe2 == 'S'
			}
		}
		if p2.Y == p1.Y+1 {
			if p2.X == p1.X {
				return pipe2 == '|' || pipe2 == 'L' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
	case 'F':
		if p2.Y == p1.Y {
			if p2.X == p1.X+1 {
				return pipe2 == '-' || pipe2 == '7' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
		if p2.Y == p1.Y+1 {
			if p2.X == p1.X {
				return pipe2 == '|' || pipe2 == 'L' || pipe2 == 'J' || pipe2 == 'S'
			}
		}
	}
	return false
}

func (p Pipeline) GetConnectingPipes(point helper.Point) []helper.Point {
	adjacentPoints := p.Grid.GetAdjacentPoints(point)
	connectedPoints := make([]helper.Point, 0)
	for _, adj := range adjacentPoints {
		if p.AreConnected(point, adj) {
			connectedPoints = append(connectedPoints, adj)
		}
	}
	return connectedPoints
}

func (p Pipeline) FindNextPipe(current, previous helper.Point) helper.Point {
	// fmt.Println(p.GetConnectingPipes(current))
	for _, connectedPoint := range p.GetConnectingPipes(current) {
		if connectedPoint != previous {
			return connectedPoint
		}
	}
	panic(fmt.Sprintf("Cannot find connected pipe for %c at (%d %d)", p.Grid.GetCharAt(current), current.X, current.Y))
	// return helper.Point{}
}

func LoadGrid(inputFile string) Pipeline {
	lines := helper.ReadInputFile(inputFile)
	result := Pipeline{Grid: helper.Grid{Rows: lines}}
	return result
}

func SolvePart1(inputFile string) string {
	pipeline := LoadGrid(inputFile)

	for _, line := range pipeline.Grid.Rows {
		fmt.Println(line)
	}

	startingPoint := pipeline.Grid.FindLocation('S')
	fmt.Printf("Starting point is %v \n", startingPoint)

	connectingPipes := pipeline.GetConnectingPipes(startingPoint)
	fmt.Println(connectingPipes)

	secondPipe := pipeline.FindNextPipe(startingPoint, helper.Point{})
	fmt.Println("Next pipe: ", secondPipe)

	previous := startingPoint
	current := secondPipe
	for routeLength := 1; ; routeLength++ {
		next := pipeline.FindNextPipe(current, previous)
		if next == startingPoint {
			return fmt.Sprint(int(routeLength/2) + 1)
		} else {
			previous = current
			current = next
			fmt.Printf("%c > ", pipeline.Grid.GetCharAt(next))
		}
		// if routeLength > 150 {
		// 	panic("Taking too long!")
		// }
	}
}

func SolvePart2(inputFile string) string {
	pipeline := LoadGrid(inputFile)

	for _, line := range pipeline.Grid.Rows {
		fmt.Println(line)
	}

	visited := make([]helper.Point, 0)

	startingPoint := pipeline.Grid.FindLocation('S')
	secondPipe := pipeline.FindNextPipe(startingPoint, helper.Point{})

	visited = append(visited, startingPoint)
	visited = append(visited, secondPipe)

	previous := startingPoint
	current := secondPipe
	for routeLength := 1; ; routeLength++ {
		visited = append(visited, current)
		next := pipeline.FindNextPipe(current, previous)
		if next == startingPoint {
			break
		} else {
			previous = current
			current = next
		}
	}

	pipeline.Grid.SetCharAt(startingPoint, '#')
	for _, p := range visited {
		pipeline.Grid.SetCharAt(p, '#')
	}

	for y := 0; y < len(pipeline.Grid.Rows); y++ {
		for x := 0; x < len(pipeline.Grid.Rows[0]); x++ {
			if pipeline.Grid.GetCharAtXY(x, y) != 'O' {
				pipeline.FillGrid(x, y)
			}
		}
	}

	fmt.Println("")
	pipeline.Grid.Draw()

	result := 0
	for y := 0; y < len(pipeline.Grid.Rows); y++ {
		for x := 0; x < len(pipeline.Grid.Rows[0]); x++ {
			if pipeline.Grid.GetCharAtXY(x, y) == 'I' {
				result++
			}
		}
	}

	return fmt.Sprint(result)
}

// Start with a point x, y
// Check if the point is an empty ground, '.'. Skip if not.
// create a set of currentPoints and add our first point to it
//  For all points in currentPoints, find their uniqueNeighbours
//  If all uniqueNeighbours are our rope, mark each point in currentPoints as "in"
//  Else add all ground points (".") from uniqueNeighbours to currentPoints
//  Repeat until either all neighbours are our rope, or until there is a "." at edge of map
//  FIXME this produces false positives when a point is surrounded by our loop but isn't within it

func (pipeline *Pipeline) FillGrid(x, y int) {

	if pipeline.Grid.GetCharAtXY(x, y) != '.' {
		return
	}

	currentPoints := map[helper.Point]bool{}
	currentPoints[helper.Point{X: x, Y: y}] = true

	for {
		uniqueNeighbours := map[helper.Point]bool{}
		for point := range currentPoints {
			for _, adjacent := range pipeline.Grid.GetAdjacents(point, false) {
				if _, ok := currentPoints[adjacent]; !ok {
					uniqueNeighbours[adjacent] = true
				}
			}
		}

		allRopes := true
		groundNeighbours := map[helper.Point]bool{}
		for neighbourPoint := range uniqueNeighbours {
			charAt := pipeline.Grid.GetCharAt(neighbourPoint)
			if charAt != '#' {
				allRopes = false
			}
			if charAt == '.' {
				groundNeighbours[neighbourPoint] = true
			}
		}
		if allRopes {
			for point := range currentPoints {
				pipeline.Grid.SetCharAt(point, 'I')
			}
			return
		} else {
			if len(groundNeighbours) == 0 {
				for point := range currentPoints {
					pipeline.Grid.SetCharAt(point, 'O')
				}
				return
			}
			for point := range groundNeighbours {
				currentPoints[point] = true
				if point.X == 0 || point.Y == 0 || point.X == pipeline.Grid.LentghX()-1 || point.Y == pipeline.Grid.Lentgh()-1 {
					// reached edge of a map, so all points are out
					for point := range currentPoints {
						pipeline.Grid.SetCharAt(point, 'O')
					}
					return
				}
			}
		}
	}
}
