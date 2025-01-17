package solution

import (
	"reflect"
	"testing"
)

func Test_smallestRange(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: [][]int{
					{4, 10, 15, 24, 26},
					{0, 9, 12, 20},
					{5, 18, 22, 30},
				},
			},
			want: []int{20, 24},
		},
		{
			name: "",
			args: args{
				nums: [][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
