package day4

import (
	"reflect"
	"testing"
)

func TestReadScratchcard(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Scratchcard
	}{
		{
			"Load simple scratchcard",
			args{"Card 1: 1 2 3 | 2 3 4"},
			Scratchcard{Id: 1, WinningNumbers: []int{1, 2, 3}, Numbers: []int{2, 3, 4}, Copies: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadScratchcard(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadScratchcard() = %v, want %v", got, tt.want)
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
			"13",
		},
		{
			"Solve real input for part 1 (regression test)",
			args{"input/input1.txt"},
			"25010",
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

func TestScratchcard_CountWinningNumbers(t *testing.T) {
	type fields struct {
		Id             int
		WinningNumbers []int
		Numbers        []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Count winning fields, simple example",
			fields{
				Id:             1,
				WinningNumbers: []int{1, 2, 3},
				Numbers:        []int{3, 2, 1},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scratchcard := &Scratchcard{
				Id:             tt.fields.Id,
				WinningNumbers: tt.fields.WinningNumbers,
				Numbers:        tt.fields.Numbers,
			}
			if got := scratchcard.CountWinningNumbers(); got != tt.want {
				t.Errorf("Scratchcard.CountWinningNumbers() = %v, want %v", got, tt.want)
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
			"30",
		},
		{
			"Solve real input for part 2 (regression test)",
			args{"input/input1.txt"},
			"9924412",
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
