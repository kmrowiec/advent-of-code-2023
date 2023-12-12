package day11

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
			"374",
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

func TestSumOfShortestPaths(t *testing.T) {
	type args struct {
		inputFile string
		multplier int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Solve example 1 for part 1",
			args{"input/example1.txt", 2},
			"374",
		},
		{
			"Solve example 1 for part 2 - expanding 10 times",
			args{"input/example1.txt", 10},
			"1030",
		},
		{
			"Solve example 1 for part 2 - expanding 100 times",
			args{"input/example1.txt", 100},
			"8410",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfShortestPaths(tt.args.inputFile, tt.args.multplier); got != tt.want {
				t.Errorf("SumOfShortestPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
