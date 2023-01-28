package ways_to_make_a_fair_array

import (
	"testing"
)

func Test_waysToMakeFair(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 1, 6, 4},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := waysToMakeFair(tt.args.nums); got != tt.want {
				t.Errorf("waysToMakeFair() = %v, want %v", got, tt.want)
			}
		})
	}
}
