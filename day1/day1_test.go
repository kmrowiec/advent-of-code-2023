package day1

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
			"Check against basic example",
			args{"input/example1.txt"},
			"142",
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
			"Check against basic example for part 2",
			args{"input/example2.txt"},
			"281",
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

func TestFindOnlyNumbers(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Replace one with 1",
			args{"1one"},
			"11",
		},
		{
			"Replace smthoneoneonenope with smth111nope",
			args{"smthoneoneonenope"},
			"111",
		},
		{
			"Replace onetwothreefourfivesixseveneightnine with 123456789",
			args{"onetwothreefourfivesixseveneightnine"},
			"123456789",
		},
		{
			"Replace oneightwoneight with 18218",
			args{"oneightwoneight"},
			"18218",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOnlyNumbers(tt.args.line); got != tt.want {
				t.Errorf("ReplaceSpelledNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
