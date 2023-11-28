// 862. 和至少为 K 的最短子数组
// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
package solution

import "testing"

func Test_shortestSubarray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{17, 85, 93, -45, -21},
				k:    150,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2},
				k:    4,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums: []int{2, -1, 2},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSubarray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("shortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
