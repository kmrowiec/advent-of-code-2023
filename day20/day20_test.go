package day20

import "testing"

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
			"Solve main example 1 for part 1",
			args{"input/example1.txt"},
			"32000000",
		},
		{
			"Solve main example 2 for part 1",
			args{"input/example2.txt"},
			"11687500",
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
