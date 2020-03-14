package longest_increasing_subsequence

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 15, 3, 7, 8, 6, 18}}, 5},
		{"", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, 6},
		{"", args{[]int{}}, 0},
		{"", args{[]int{0}}, 1},
		{"", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
