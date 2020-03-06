package besttimetobuyandsellstockwithcooldown

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
		// {
		// 	"[1,2,1,2,1,3,5,1]",
		// 	args{
		// 		[]int{1, 2, 1, 2, 1, 3, 5, 1},
		// 	},
		// 	5,
		// },
		// {
		// 	"[2,1]",
		// 	args{
		// 		[]int{2, 1},
		// 	},
		// 	0,
		// },
		{
			"[1,2,3,0,2]",
			args{
				[]int{1, 2, 3, 0, 2},
			},
			3,
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
