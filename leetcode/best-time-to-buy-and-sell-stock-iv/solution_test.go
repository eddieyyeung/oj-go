package besttimetobuyandsellstockiv

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
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
				2,
				[]int{2, 4, 1},
			},
			2,
		},
		{
			"case2",
			args{
				2,
				[]int{3, 2, 6, 5, 0, 3},
			},
			7,
		},
		{
			"case3",
			args{
				4,
				[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			},
			15,
		},
		{
			"case4",
			args{
				2,
				[]int{2, 1, 4},
			},
			3,
		},
		{
			"case5",
			args{
				2,
				[]int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
