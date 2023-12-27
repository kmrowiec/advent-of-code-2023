package day21

import "testing"

func TestSimulate(t *testing.T) {
	type args struct {
		inputFile string
		steps     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Simulate the main example",
			args{"input/example1.txt", 6},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Simulate(tt.args.inputFile, tt.args.steps); got != tt.want {
				t.Errorf("Simulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
