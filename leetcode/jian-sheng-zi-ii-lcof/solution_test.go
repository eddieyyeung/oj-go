package jianshengziiilcof

import "testing"

func Test_cuttingRope(t *testing.T) {
	type args struct {
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
				n: 120,
			},
			want: 953271190,
		},
		{
			name: "",
			args: args{
				n: 4,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				n: 10,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
