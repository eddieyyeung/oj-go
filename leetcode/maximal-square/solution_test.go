package maximal_square

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{matrix: [][]byte{
			{'1', '0', '1', '1', '0', '1'},
			{'1', '1', '1', '1', '1', '1'},
			{'0', '1', '1', '0', '1', '1'},
			{'1', '1', '1', '0', '1', '0'},
			{'0', '1', '1', '1', '1', '1'},
			{'1', '1', '0', '1', '1', '1'},
		}}, 4},
		{"", args{matrix: [][]byte{
			{'0', '0', '0'},
			{'1', '0', '0'},
			{'0', '0', '0'},
		}}, 1},
		{"", args{matrix: [][]byte{
			{'1', '1', '0', '1'},
			{'1', '1', '0', '1'},
			{'1', '1', '1', '1'},
		}}, 4},
		{"", args{matrix: [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
