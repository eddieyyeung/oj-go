package shortestpathwithalternatingcolors

import (
	"reflect"
	"testing"
)

func Test_shortestAlternatingPaths(t *testing.T) {
	type args struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				n:         3,
				redEdges:  [][]int{{0, 1}},
				blueEdges: [][]int{{1, 2}},
			},
			want: []int{0, 1, 2},
		},
		{
			args: args{
				n:         5,
				redEdges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
				blueEdges: [][]int{{1, 2}, {2, 3}, {3, 1}},
			},
			want: []int{0, 1, 2, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := shortestAlternatingPaths(tt.args.n, tt.args.redEdges, tt.args.blueEdges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestAlternatingPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
