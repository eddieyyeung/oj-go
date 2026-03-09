package climbing_stairs_ii

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n     int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:     4,
				costs: []int{1, 2, 3, 4},
			},
			want: 13,
		},
		{
			name: "",
			args: args{
				n:     4,
				costs: []int{5, 1, 6, 2},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n, tt.args.costs); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
