package checkknighttourconfiguration

import "testing"

func Test_checkValidGrid(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				grid: [][]int{{24, 11, 22, 17, 4}, {21, 16, 5, 12, 9}, {6, 23, 10, 3, 18}, {15, 20, 1, 8, 13}, {0, 7, 14, 19, 2}},
			},
			want: false,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		grid: [][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}},
		// 	},
		// 	want: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidGrid(tt.args.grid); got != tt.want {
				t.Errorf("checkValidGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
