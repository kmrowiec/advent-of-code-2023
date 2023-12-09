package day7

import (
	"reflect"
	"testing"
)

func TestReadHand(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Hand
	}{
		{
			"Read a hand - high card",
			args{"23456 123"},
			Hand{Raw: "23456", Strength: 6, Bet: 123},
		},
		{
			"Read a hand - pair",
			args{"23256 123"},
			Hand{Raw: "23256", Strength: 5, Bet: 123},
		},
		{
			"Read a hand - two pair",
			args{"23253 123"},
			Hand{Raw: "23253", Strength: 4, Bet: 123},
		},
		{
			"Read a hand - three of a kind",
			args{"23252 123"},
			Hand{Raw: "23252", Strength: 3, Bet: 123},
		},
		{
			"Read a hand - full house",
			args{"23232 123"},
			Hand{Raw: "23232", Strength: 2, Bet: 123},
		},
		{
			"Read a hand - four of a kind",
			args{"AA8AA 123"},
			Hand{Raw: "AA8AA", Strength: 1, Bet: 123},
		},
		{
			"Read a hand - five of a kind",
			args{"AAAAA 123"},
			Hand{Raw: "AAAAA", Strength: 0, Bet: 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadHand(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHand() = %v, want %v", got, tt.want)
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
			"6440",
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

func TestCompareRawHands(t *testing.T) {
	type args struct {
		first  Hand
		second Hand
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Compare hands 1",
			args{first: ReadHand("KK677 28"),
				second: ReadHand("KTJJT 220")},
			false,
		},
		{
			"Compare hands 2",
			args{first: ReadHand("T55J5 684"),
				second: ReadHand("QQQJA 483")},
			true,
		},
		{
			"Compare hands 3a",
			args{first: ReadHand("77767 684"),
				second: ReadHand("66656 483")},
			false,
		},
		{
			"Compare hands 3b",
			args{first: ReadHand("66656 483"),
				second: ReadHand("77767 684")},
			true,
		},
		{
			"Compare hands 4",
			args{first: ReadHand("KKKQK 483"),
				second: ReadHand("KKKKQ 684")},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareRawHands(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CompareRawHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadHandWithJokers(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Hand
	}{
		{
			"Read a hand with 1 joker - high card",
			args{"2345J 22"},
			Hand{Raw: "2345J", Strength: 5, Bet: 22},
		},
		{
			"Read a hand with 1 joker - pair",
			args{"2245J 22"},
			Hand{Raw: "2245J", Strength: 3, Bet: 22},
		},
		{
			"Read a hand with 1 joker - 2 pair",
			args{"2323J 22"},
			Hand{Raw: "2323J", Strength: 2, Bet: 22},
		},
		{
			"Read a hand with 1 joker - three",
			args{"2322J 22"},
			Hand{Raw: "2322J", Strength: 1, Bet: 22},
		},
		{
			"Read a hand with 1 joker - four",
			args{"2222J 22"},
			Hand{Raw: "2222J", Strength: 0, Bet: 22},
		},
		{
			"Read a hand with 2 jokers - pair",
			args{"234JJ 22"},
			Hand{Raw: "234JJ", Strength: 3, Bet: 22},
		},
		{
			"Read a hand with 2 jokers - 2 pair",
			args{"334JJ 22"},
			Hand{Raw: "334JJ", Strength: 1, Bet: 22},
		},
		{
			"Read a hand with 3 jokers - three",
			args{"23JJJ 22"},
			Hand{Raw: "23JJJ", Strength: 1, Bet: 22},
		},
		{
			"Read a hand with 3 jokers - full",
			args{"33JJJ 22"},
			Hand{Raw: "33JJJ", Strength: 0, Bet: 22},
		},
		{
			"Read a hand with 4 jokers - four",
			args{"3JJJJ 22"},
			Hand{Raw: "3JJJJ", Strength: 0, Bet: 22},
		},
		{
			"Read a hand with a joker 2",
			args{"QQQJA 483"},
			Hand{Raw: "QQQJA", Strength: 1, Bet: 483},
		},
		{
			"Read a hand with a joker 3",
			args{"KK677 483"},
			Hand{Raw: "KK677", Strength: 4, Bet: 483},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadHandWithJokers(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHandWithJokers() = %v, want %v", got, tt.want)
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
			"5905",
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
