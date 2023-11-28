package solution

import "testing"

func Test_findMaxValueOfEquation(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				points: [][]int{{-17, 5}, {-10, -8}, {-5, -13}, {-2, 7}, {8, -14}},
				k:      4,
			},
			want: -3,
		},
		{
			name: "",
			args: args{
				points: [][]int{{-19, -12}, {-13, -18}, {-12, 18}, {-11, -8}, {-8, 2}, {-7, 12}, {-5, 16}, {-3, 9}, {1, -7}, {5, -4}, {6, -20}, {10, 4}, {16, 4}, {19, -9}, {20, 19}},
				k:      6,
			},
			want: 35,
		},
		{
			name: "",
			args: args{
				points: [][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}},
				k:      1,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				points: [][]int{{0, 0}, {3, 0}, {9, 2}},
				k:      3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxValueOfEquation(tt.args.points, tt.args.k); got != tt.want {
				t.Errorf("findMaxValueOfEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
