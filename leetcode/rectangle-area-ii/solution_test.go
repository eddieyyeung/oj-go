package rectangle_area_ii

import (
	"fmt"
	"testing"
)

func Test_rectangleArea(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				rectangles: [][]int{
					{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1},
				},
			},
			want: 6,
		},
		{
			args: args{
				rectangles: [][]int{
					{0, 0, 1000000000, 1000000000},
				},
			},
			want: 49,
		},
		{
			args: args{
				rectangles: [][]int{
					{0, 0, 2, 2},
					{1, 1, 3, 3},
				},
			},
			want: 7,
		},
		{
			args: args{
				rectangles: [][]int{
					{49, 40, 62, 100}, {11, 83, 31, 99}, {19, 39, 30, 99},
				},
			},
			want: 1584,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := rectangleArea(tt.args.rectangles); got != tt.want {
				t.Errorf("rectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
