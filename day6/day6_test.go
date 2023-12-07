package day6

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
			"Solve example for part 1",
			args{"input/example1.txt"},
			"288",
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

func TestLoadRaceRecords(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []RaceRecord
	}{
		{
			"Read example input as RaceRecords",
			args{[]string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			}},
			[]RaceRecord{
				{Time: 7, Distance: 9},
				{Time: 15, Distance: 40},
				{Time: 30, Distance: 200},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadRaceRecords(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRaceRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRaceRecord_GetMarginOfError(t *testing.T) {
	type fields struct {
		Time     int
		Distance int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Calculate margin of error - simple example 1",
			fields{
				Time:     7,
				Distance: 9,
			},
			4,
		},
		{
			"Calculate margin of error - simple example 2",
			fields{
				Time:     15,
				Distance: 40,
			},
			8,
		},
		{
			"Calculate margin of error - simple example 3",
			fields{
				Time:     30,
				Distance: 200,
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raceRecord := &RaceRecord{
				Time:     tt.fields.Time,
				Distance: tt.fields.Distance,
			}
			if got := raceRecord.GetMarginOfError(); got != tt.want {
				t.Errorf("RaceRecord.GetMarginOfError() = %v, want %v", got, tt.want)
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
			"Solve example for part 1",
			args{"input/example1.txt"},
			"71503",
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

func TestLoadRaceRecordsPartTwo(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want RaceRecord
	}{
		{
			"Read example input as RaceRecords",
			args{[]string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			}},
			RaceRecord{Time: 71530, Distance: 940200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadRaceRecordsPartTwo(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRaceRecordsPartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
