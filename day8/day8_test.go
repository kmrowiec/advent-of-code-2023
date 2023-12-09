package day8

import (
	"reflect"
	"testing"
)

func TestReadNode(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 Node
	}{
		{
			"Read simple node",
			args{"AAA = (BBB, CCC)"},
			"AAA",
			Node{Left: "BBB", Right: "CCC"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ReadNode(tt.args.line)
			if got != tt.want {
				t.Errorf("ReadNode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadNode() got1 = %v, want %v", got1, tt.want1)
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
			"Solve example 1 for part 1",
			args{"input/example1.txt"},
			"2",
		},
		{
			"Solve example 2 for part 1",
			args{"input/example2.txt"},
			"6",
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
			"Solve example 3 for part 1",
			args{"input/example3.txt"},
			"6",
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
