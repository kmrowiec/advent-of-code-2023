package day3

import (
	"reflect"
	"testing"

	"dev.kmrowiec/aoc/helper"
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
			"Solve example 1",
			args{"input/example1.txt"},
			"4361",
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

func TestFindNumbersInLine(t *testing.T) {
	type args struct {
		grid helper.Grid
		y    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Find all valid numbers in a grid row (1)",
			args{
				grid: LoadGrid("input/example1.txt"),
				y:    0,
			},
			[]int{467},
		},
		{
			"Find all valid numbers in a grid row (2)",
			args{
				grid: LoadGrid("input/example1.txt"),
				y:    9,
			},
			[]int{664, 598},
		},
		{
			"Find all valid numbers in a grid row (2)",
			args{
				grid: LoadGrid("input/example1.txt"),
				y:    3,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumbersInLine(tt.args.grid, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNumbersInLine() = %v, want %v", got, tt.want)
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
			"467835",
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
