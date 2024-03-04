package solution

import "testing"

func Test_maxTaxiEarnings(t *testing.T) {
	type args struct {
		n     int
		rides [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				n:     10,
				rides: [][]int{{2, 3, 6}, {8, 9, 8}, {5, 9, 7}, {8, 9, 1}, {2, 9, 2}, {9, 10, 6}, {7, 10, 10}, {6, 7, 9}, {4, 9, 7}, {2, 3, 1}},
			},
			want: 33,
		},
		{
			name: "",
			args: args{
				n: 5,
				rides: [][]int{
					{2, 5, 4},
					{1, 5, 1},
				},
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				n:     20,
				rides: [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTaxiEarnings(tt.args.n, tt.args.rides); got != tt.want {
				t.Errorf("maxTaxiEarnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
