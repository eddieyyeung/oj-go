package longest_continuous_increasing_subsequence

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 3, 5, 4, 7}}, 3},
		{"", args{[]int{2, 2, 2, 2, 2}}, 1},
		{"", args{[]int{1, 3, 5, 7}}, 4},
		{"", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
