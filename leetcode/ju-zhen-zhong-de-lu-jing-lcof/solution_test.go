package juzhenzhongdelujinglcof

import "testing"

func Test_wordPuzzle(t *testing.T) {
	type args struct {
		grid   [][]byte
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
				},
				target: "ABCCED",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
				},
				target: "SEE",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				grid: [][]byte{
					{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
				},
				target: "ABCB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPuzzle(tt.args.grid, tt.args.target); got != tt.want {
				t.Errorf("wordPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
