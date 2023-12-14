package day13

import (
	"testing"

	"dev.kmrowiec/aoc/helper/grid"
)

func TestSolvePart1(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Solve example 1 for part 1",
			args{"input/example1.txt"},
			"405",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart1(tt.args.inputFile); got != tt.want {
				t.Errorf("SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflect(t *testing.T) {
	type args struct {
		grid *grid.Grid
	}
	example2Grid := grid.GridFromFile("input/example2.txt")
	example3Grid := grid.GridFromFile("input/example3.txt")
	tests := []struct {
		name              string
		args              args
		wantVertical      bool
		wantReflectedRows int
	}{
		{
			"Find reflection in a grid - edge case 1",
			args{grid: &example2Grid},
			false,
			1,
		},
		{
			"Find reflection in a grid - edge case 2",
			args{grid: &example3Grid},
			true,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVertical, gotReflectedRows := Reflect(tt.args.grid)
			if gotVertical != tt.wantVertical {
				t.Errorf("Reflect() gotVertical = %v, want %v", gotVertical, tt.wantVertical)
			}
			if gotReflectedRows != tt.wantReflectedRows {
				t.Errorf("Reflect() gotReflectedRows = %v, want %v", gotReflectedRows, tt.wantReflectedRows)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Solve example 1 for part 2",
			args{"input/example1.txt"},
			"400",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart2(tt.args.inputFile); got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindVariations(t *testing.T) {
	type args struct {
		g grid.Grid
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test find variations",
			args{g: grid.GridFromLines([]string{".#....###",
				"#.#######",
				"#.#######",
				".#....###",
				"..##.#...",
				".###.#.##",
				".#...#.##",
				"..####.##",
				"...##..#."})},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindVariations(tt.args.g); len(got) != tt.want {
				t.Errorf("FindVariations() = %v, want %v", got, tt.want)
			}
		})
	}
}
