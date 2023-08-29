package superpow

import "testing"

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				a: 2,
				b: []int{3},
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				a: 2,
				b: []int{1, 0},
			},
			want: 1024,
		},
		{
			name: "",
			args: args{
				a: 2147483647,
				b: []int{2, 0, 0},
			},
			want: 1198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
