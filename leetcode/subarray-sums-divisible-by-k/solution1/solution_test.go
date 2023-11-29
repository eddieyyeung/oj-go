// 974. 和可被 K 整除的子数组
// https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/
package solution

import "testing"

func Test_subarraysDivByK(t *testing.T) {
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
				nums: []int{-5, 12},
				k:    3,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{-1, 2, 9},
				k:    2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{4, 5, 0, -2, -3, 1},
				k:    5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
