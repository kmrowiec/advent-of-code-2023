package day16

import (
	"fmt"

	"dev.kmrowiec/aoc/helper/grid"
	"github.com/schollz/progressbar/v3"
)

type Solver struct{}

func (s *Solver) PartOne() string {
	return SolvePart1("day16/input/input1.txt")
}

func (s *Solver) PartTwo() string {
	return SolvePart2("day16/input/input1.txt")
}

func SolvePart1(inputFile string) string {
	g := grid.GridFromFile(inputFile)
	energised := map[grid.Point]map[rune]bool{}
	SimulateBeams(&g, energised, []Beam{{X: 0, Y: 0, Direction: 'E'}})
	return fmt.Sprint(len(energised))
}

func SolvePart2(inputFile string) string {
	g := grid.GridFromFile(inputFile)
	startingBeams := GenerateStartingBeams(&g)
	bar := progressbar.Default(int64(len(startingBeams)))
	bestResult := 0
	for _, beam := range startingBeams {
		energised := map[grid.Point]map[rune]bool{}
		SimulateBeams(&g, energised, []Beam{beam})
		if result := len(energised); result > bestResult {
			bestResult = result
		}
		bar.Add(1)
	}
	return fmt.Sprint(bestResult)
}

type Beam struct {
	X, Y      int
	Direction rune
}

func (b *Beam) String() string {
	return fmt.Sprintf("Beam at %d, %d going %c", b.X, b.Y, b.Direction)
}

func SimulateBeams(g *grid.Grid, e map[grid.Point]map[rune]bool, beams []Beam) {
	for len(beams) > 0 {
		beams = SimulateStep(g, e, beams)
	}
}

func SimulateStep(g *grid.Grid, e map[grid.Point]map[rune]bool, beams []Beam) []Beam {
	newBeams := make([]Beam, 0)
	for _, beam := range beams {
		// If the beam reached outside of the grid, it's ignored
		if IsOutside(g, &beam) {
			continue
		}
		// Then we check if the point was already energised by a beam with the same dirextion
		// If not, mark it as energised, otherwise ignore the beam from further simulation
		if energised, ok := e[grid.PointFromXY(beam.X, beam.Y)]; ok {
			if _, ok := energised[beam.Direction]; ok {
				continue
			} else {
				energised[beam.Direction] = true
			}
		} else {
			e[grid.PointFromXY(beam.X, beam.Y)] = map[rune]bool{beam.Direction: true}
		}
		// Then "move" the Beam to its next location
		switch g.GetCharAtXY(beam.X, beam.Y) {
		case '.':
			switch beam.Direction {
			case 'N':
				beam.Y--
			case 'S':
				beam.Y++
			case 'W':
				beam.X--
			case 'E':
				beam.X++
			}
			newBeams = append(newBeams, beam)
		case '|':
			switch beam.Direction {
			case 'N':
				beam.Y--
				newBeams = append(newBeams, beam)
			case 'S':
				beam.Y++
				newBeams = append(newBeams, beam)
			default:
				newBeams = append(newBeams, Beam{X: beam.X, Y: beam.Y - 1, Direction: 'N'})
				newBeams = append(newBeams, Beam{X: beam.X, Y: beam.Y + 1, Direction: 'S'})
			}
		case '-':
			switch beam.Direction {
			case 'W':
				beam.X--
				newBeams = append(newBeams, beam)
			case 'E':
				beam.X++
				newBeams = append(newBeams, beam)
			default:
				newBeams = append(newBeams, Beam{X: beam.X + 1, Y: beam.Y, Direction: 'E'})
				newBeams = append(newBeams, Beam{X: beam.X - 1, Y: beam.Y, Direction: 'W'})
			}
		case '\\':
			switch beam.Direction {
			case 'N':
				newBeams = append(newBeams, Beam{X: beam.X - 1, Y: beam.Y, Direction: 'W'})
			case 'S':
				newBeams = append(newBeams, Beam{X: beam.X + 1, Y: beam.Y, Direction: 'E'})
			case 'E':
				newBeams = append(newBeams, Beam{X: beam.X, Y: beam.Y + 1, Direction: 'S'})
			case 'W':
				newBeams = append(newBeams, Beam{X: beam.X, Y: beam.Y - 1, Direction: 'N'})
			}
		case '/':
			switch beam.Direction {
			case 'N':
				newBeams = append(newBeams, Beam{X: beam.X + 1, Y: beam.Y, Direction: 'E'})
			case 'S':
				newBeams = append(newBeams, Beam{X: beam.X - 1, Y: beam.Y, Direction: 'W'})
			case 'E':
				newBeams = append(newBeams, Beam{X: beam.X, Y: beam.Y - 1, Direction: 'N'})
			case 'W':
				newBeams = append(newBeams, Beam{X: beam.X, Y: beam.Y + 1, Direction: 'S'})
			}
		}
	}
	return newBeams
}

func IsOutside(g *grid.Grid, beam *Beam) bool {
	return beam.X < 0 || beam.Y < 0 || beam.X >= g.RowLength() || beam.Y >= g.ColumnLength()
}

func GenerateStartingBeams(g *grid.Grid) []Beam {
	result := make([]Beam, 0)
	for x := 0; x < g.RowLength(); x++ {
		result = append(result, Beam{X: x, Y: 0, Direction: 'S'})
	}
	for x := 0; x < g.RowLength(); x++ {
		result = append(result, Beam{X: x, Y: g.ColumnLength() - 1, Direction: 'N'})
	}
	for y := 0; y < g.ColumnLength(); y++ {
		result = append(result, Beam{X: 0, Y: y, Direction: 'E'})
	}
	for y := 0; y < g.ColumnLength(); y++ {
		result = append(result, Beam{X: g.RowLength() - 1, Y: y, Direction: 'W'})
	}
	return result
}
