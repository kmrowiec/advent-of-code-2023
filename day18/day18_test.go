package day18

import (
	"reflect"
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
			"Solve main example for part 1",
			args{"input/example1.txt"},
			"62",
		},
		{
			"Solve a very simple example for part 1",
			args{"input/example2.txt"},
			"25",
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
			"Solve main example for part 2",
			args{"input/example1.txt"},
			"952408144115",
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

func TestDecodeInstruction(t *testing.T) {
	type args struct {
		hexInput string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Decode an example instruction",
			args{"#70c710"},
			[]string{"R", "461937"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeInstruction(tt.args.hexInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
