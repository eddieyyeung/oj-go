package maximum_multiplication_score

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				a: []int{3, 2, 5, 6},
				b: []int{2, -6, 4, -5, -3, 2, -7},
			},
			want: 26,
		},
		{
			name: "",
			args: args{
				a: []int{-1, 4, 5, -2},
				b: []int{-5, -1, -3, -2, -4},
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				a: []int{100000, 100000, 100000, 100000},
				b: []int{-100000, -100000, -100000, -100000},
			},
			want: -40000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
