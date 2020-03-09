package besttimetobuyandsellstockiii

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			6,
		},
		{
			"case2",
			args{
				[]int{1, 2, 3, 4, 5},
			},
			4,
		},
		{
			"case3",
			args{
				[]int{7, 6, 4, 3, 1},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
