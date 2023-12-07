package helper

import (
	"reflect"
	"testing"
)

func TestGrid_GetCharAt(t *testing.T) {
	type fields struct {
		rows []string
	}
	type args struct {
		point Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rune
	}{
		{
			"Get an item from a grid",
			fields{
				[]string{"123", "456", "789"},
			},
			args{
				point: Point{0, 0},
			},
			'1',
		},
		{
			"Get an item from a grid",
			fields{
				[]string{"123", "456", "789"},
			},
			args{
				point: Point{2, 2},
			},
			'9',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := &Grid{
				Rows: tt.fields.rows,
			}
			if got := grid.GetCharAt(tt.args.point); got != tt.want {
				t.Errorf("Grid.GetCharAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_GetAdjacentPoints(t *testing.T) {
	type fields struct {
		Rows []string
	}
	type args struct {
		point Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Point
	}{
		{
			"Get adjacent points - easy case",
			fields{
				[]string{"123", "456", "789"},
			},
			args{
				point: Point{1, 1},
			},
			[]Point{{0, 1}, {0, 0}, {0, 2}, {1, 0}, {1, 2}, {2, 1}, {2, 0}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := &Grid{
				Rows: tt.fields.Rows,
			}
			if got := grid.GetAdjacentPoints(tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.GetAdjacentPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
