package solution

import "testing"

func Test_superEggDrop(t *testing.T) {
	type args struct {
		k int
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
				k: 2,
				n: 4,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				k: 2,
				n: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				k: 2,
				n: 2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				k: 1,
				n: 2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				k: 2,
				n: 6,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				k: 3,
				n: 14,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
