// 2952. 需要添加的硬币的最小数量
// https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/description/
package solution

import "testing"

func Test_minimumAddedCoins(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins:  []int{1, 4, 10},
				target: 20,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				coins:  []int{1, 4, 10, 5, 7, 19},
				target: 19,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				coins:  []int{1, 1, 1},
				target: 20,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAddedCoins(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("minimumAddedCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
