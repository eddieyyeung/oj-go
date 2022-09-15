package special_array_with_x_elements_greater_than_or_equal_x

import (
	"fmt"
	"testing"
)

func Test_specialArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 5},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 0},
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{0, 4, 3, 0, 4},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{3, 6, 7, 7, 0},
			},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := specialArray(tt.args.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
