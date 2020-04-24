package merge_intervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"", args{[][]int{{1, 4}, {2, 3}}}, [][]int{{1, 4}}},
		{"", args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"", args{[][]int{{1, 4}, {4, 5}}}, [][]int{{1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
