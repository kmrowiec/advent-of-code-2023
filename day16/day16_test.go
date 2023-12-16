package day16

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
			"Solve example for part 1",
			args{"input/example1.txt"},
			"46",
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
			"Solve example for part 2",
			args{"input/example1.txt"},
			"51",
		},
		{
			"Solve input for part 2",
			args{"input/input1.txt"},
			"7932",
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
