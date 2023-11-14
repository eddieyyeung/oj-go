package coinchange

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins:  []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422},
				amount: 9864,
			},
			want: 24,
		},
		{
			"[1, 2, 5]",
			args{
				[]int{1, 2, 5},
				11,
			},
			3,
		},
		{
			"[2]",
			args{
				[]int{2},
				3,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
