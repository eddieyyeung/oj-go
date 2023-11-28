// 2944. 购买水果需要的最少金币数
// https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/description/
package solution

import "testing"

func Test_minimumCoins(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				prices: []int{3, 1, 2},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				prices: []int{1, 10, 1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCoins(tt.args.prices); got != tt.want {
				t.Errorf("minimumCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
