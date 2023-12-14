package day12

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
			"Solve example 1 for part 1",
			args{"input/example1.txt"},
			"21",
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

func TestIsMatching(t *testing.T) {
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test matching patterns 1",
			args{"#.#.###", "1,1,3"},
			true,
		},
		{
			"Test matching patterns 2",
			args{"#.#.#.#", "1,1,3"},
			false,
		},
		{
			"Test matching patterns 3",
			args{"...#.######..#####.", "1,6,5"},
			true,
		},
		{
			"Test matching patterns 4",
			args{"..#..######..#####.", "1,6,5"},
			true,
		},
		{
			"Test matching patterns 5",
			args{".##..######..#####.", "1,6,5"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatching(tt.args.pattern, tt.args.target); got != tt.want {
				t.Errorf("IsMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateOptions(t *testing.T) {
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Generate options 1",
			args{"#.#.###", "1,1,3"},
			[]string{"#.#.###"},
		},
		{
			"Generate options 2",
			args{"???.###", "1,1,3"},
			[]string{"#.#.###"},
		},
		{
			"Generate options 3",
			args{".??..??...?##.", "1,1,3"},
			[]string{"..#...#...###.", "..#..#....###.", ".#....#...###.", ".#...#....###."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateOptions(tt.args.pattern, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountOptions(t *testing.T) {
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Count options 1",
			args{".??..??...?##.", "1,1,3"},
			4,
		},
		{
			"Count options 2",
			args{"?#?#?#?#?#?#?#?", "1,3,1,6"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOptions(tt.args.pattern, tt.args.target); got != tt.want {
				t.Errorf("CountOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnfold(t *testing.T) {
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name           string
		args           args
		wantNewPattern string
		wantNewTarget  string
	}{
		{
			"Unfold example - 1",
			args{"???.###", "1,1,3"},
			"???.###????.###????.###????.###????.###",
			"1,1,3,1,1,3,1,1,3,1,1,3,1,1,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewPattern, gotNewTarget := Unfold(tt.args.pattern, tt.args.target)
			if gotNewPattern != tt.wantNewPattern {
				t.Errorf("Unfold() gotNewPattern = %v, want %v", gotNewPattern, tt.wantNewPattern)
			}
			if gotNewTarget != tt.wantNewTarget {
				t.Errorf("Unfold() gotNewTarget = %v, want %v", gotNewTarget, tt.wantNewTarget)
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
			"525152",
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

func TestCountOptionsUnfolded(t *testing.T) {
	type args struct {
		pattern string
		target  string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			"Count options unfolded 1",
			args{".??..??...?##.", "1,1,3"},
			1,
		},
		{
			"Count options unfolded 2",
			args{"?###????????", "3,2,1"},
			506250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CountOptionsUnfolded(tt.args.pattern, tt.args.target); gotResult != tt.wantResult {
				t.Errorf("CountOptionsUnfolded() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
