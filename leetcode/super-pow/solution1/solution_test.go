package solution1

import "testing"

func Test_superpow(t *testing.T) {
	type args struct {
		a int
		b int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a: 2,
				b: 10,
				n: 7,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superpow(tt.args.a, tt.args.b, tt.args.n); got != tt.want {
				t.Errorf("superpow() = %v, want %v", got, tt.want)
			}
		})
	}
}
