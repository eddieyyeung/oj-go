package remove_one_element_to_make_the_array_strictly_increasing

import (
	"fmt"
	"testing"
)

func Test_canBeIncreasing(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{1, 2, 10, 5, 7},
			},
			want: true,
		},
		{
			args: args{
				nums: []int{2, 3, 1, 2},
			},
			want: false,
		},
		{
			args: args{
				nums: []int{1, 1, 1},
			},
			want: false,
		},
		{
			args: args{
				nums: []int{105, 924, 32, 968},
			},
			want: true,
		},
		{
			args: args{
				nums: []int{100, 21, 100},
			},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := canBeIncreasing(tt.args.nums); got != tt.want {
				t.Errorf("canBeIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
