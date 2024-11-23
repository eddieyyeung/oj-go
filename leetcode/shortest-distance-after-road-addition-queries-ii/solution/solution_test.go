package solution

import (
	"reflect"
	"testing"
)

func Test_shortestDistanceAfterQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				n:       4,
				queries: [][]int{{0, 3}, {0, 2}},
			},
			want: []int{1, 1},
		},
		{
			name: "",
			args: args{
				n:       7,
				queries: [][]int{{1, 3}, {0, 3}},
			},
			want: []int{5, 4},
		},
		{
			name: "",
			args: args{
				n:       5,
				queries: [][]int{{1, 4}, {2, 4}},
			},
			want: []int{2, 2},
		},
		{
			name: "",
			args: args{
				n:       5,
				queries: [][]int{{2, 4}, {0, 2}, {0, 4}},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistanceAfterQueries(tt.args.n, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
