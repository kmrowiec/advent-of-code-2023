package day2

import (
	"reflect"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	type args struct {
		inputFile string
		criteria  map[string]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Solve example input for part 1",
			args{"input/example1.txt", map[string]int{"blue": 14, "red": 12, "green": 13}},
			"8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart1(tt.args.inputFile, tt.args.criteria); got != tt.want {
				t.Errorf("SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			"Read simplest game",
			args{"Game 1: 1 green"},
			Game{1, map[string]int{"green": 1}},
		},
		{
			"Read more complex game",
			args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			Game{1, map[string]int{"blue": 6, "red": 4, "green": 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadGame(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPossible(t *testing.T) {
	type args struct {
		game     Game
		criteria map[string]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Very simple check - false",
			args{Game{1, map[string]int{"green": 2}}, map[string]int{"green": 1}},
			false,
		},
		{
			"Very simple check - true",
			args{Game{1, map[string]int{"green": 1}}, map[string]int{"green": 2}},
			true,
		},
		{
			"Very simple check - true when equals",
			args{Game{1, map[string]int{"green": 2}}, map[string]int{"green": 2}},
			true,
		},
		{
			"More complex example 1",
			args{
				Game{1, map[string]int{"green": 3, "blue": 4, "red": 1}},
				map[string]int{"blue": 14, "red": 12, "green": 13}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPossible(tt.args.game, tt.args.criteria); got != tt.want {
				t.Errorf("IsPossible() = %v, want %v", got, tt.want)
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
			"Solve example input for part 2",
			args{"input/example1.txt"},
			"2286",
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
