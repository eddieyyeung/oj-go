package insertinterval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				intervals: [][]int{
					{3, 5},
					{12, 15},
				},
				newInterval: []int{6, 6},
			},
			want: [][]int{
				{3, 5},
				{6, 6},
				{12, 15},
			},
		},
		{
			name: "",
			args: args{
				intervals: [][]int{
					{1, 5},
				},
				newInterval: []int{0, 3},
			},
			want: [][]int{{0, 5}},
		},
		{
			name: "",
			args: args{
				intervals: [][]int{
					{1, 5},
				},
				newInterval: []int{6, 8},
			},
			want: [][]int{
				{1, 5},
				{6, 8},
			},
		},
		{
			name: "",
			args: args{
				intervals: [][]int{
					{1, 3},
					{6, 9},
				},
				newInterval: []int{2, 5},
			},
			want: [][]int{
				{1, 5},
				{6, 9},
			},
		},
		{
			name: "",
			args: args{
				intervals: [][]int{
					{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
				},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
