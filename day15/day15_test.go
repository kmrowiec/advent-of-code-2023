package day15

import (
	"testing"
)

func TestApplyHash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"HASH",
			args{"HASH"},
			52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyHash(tt.args.text); got != tt.want {
				t.Errorf("ApplyHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			"1320",
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
			"145",
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
