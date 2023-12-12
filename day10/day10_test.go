package day10

import (
	"testing"
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
			"4",
		},
		{
			"Solve example 2 for part 1",
			args{"input/example2.txt"},
			"8",
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
			"1",
		},
		{
			"Solve example 3 for part 2",
			args{"input/example3.txt"},
			"4",
		},
		{
			"Solve example 4 for part 2",
			args{"input/example4.txt"},
			"8",
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
