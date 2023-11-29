// 1590. 使数组和能被 P 整除
// https://leetcode.cn/problems/make-sum-divisible-by-p/description/
package solution

import "testing"

func Test_minSubarray(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{4, 4, 2},
				p:    7,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 1, 4, 2},
				p:    6,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{6, 3, 5, 2},
				p:    9,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3},
				p:    3,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3},
				p:    7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubarray(tt.args.nums, tt.args.p); got != tt.want {
				t.Log("args:", tt.args)
				t.Errorf("minSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
