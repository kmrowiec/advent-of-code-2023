package day5

import (
	"reflect"
	"testing"
)

func TestReadAlMap(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want AlMap
	}{
		{
			"Read simple mapping line",
			args{"1 2 3"},
			AlMap{
				start:   2,
				end:     5,
				mapping: 1,
			},
		},
		{
			"Read a random mapping line from input",
			args{"1299671458 4098616850 99419617"},
			AlMap{
				start:   4098616850,
				end:     4198036467,
				mapping: 1299671458,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadAlMap(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAlMap() = %v, want %v", got, tt.want)
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
			"35",
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
			"46",
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
