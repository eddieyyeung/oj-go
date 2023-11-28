// 2398. 预算内的最多机器人数目
// https://leetcode.cn/problems/maximum-number-of-robots-within-budget/description/
package solution

import "testing"

func Test_maximumRobots(t *testing.T) {
	type args struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				chargeTimes:  []int{3, 6, 1, 3, 4},
				runningCosts: []int{2, 1, 3, 4, 5},
				budget:       25,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				chargeTimes:  []int{11, 12, 19},
				runningCosts: []int{10, 8, 7},
				budget:       19,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRobots(tt.args.chargeTimes, tt.args.runningCosts, tt.args.budget); got != tt.want {
				t.Errorf("maximumRobots() = %v, want %v", got, tt.want)
			}
		})
	}
}
