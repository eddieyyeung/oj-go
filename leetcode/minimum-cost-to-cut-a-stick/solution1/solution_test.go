// https://leetcode.cn/problems/minimum-cost-to-cut-a-stick
package minimumcosttocutastick

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		n    int
		cuts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:    30,
				cuts: []int{13, 25, 16, 20, 26, 5, 27, 8, 23, 14, 6, 15, 21, 24, 29, 1, 19, 9, 3},
			},
			want: 0,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		n:    7,
		// 		cuts: []int{1, 3, 4, 5},
		// 	},
		// 	want: 16,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		n:    9,
		// 		cuts: []int{5, 6, 1, 4, 2},
		// 	},
		// 	want: 22,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.n, tt.args.cuts); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
