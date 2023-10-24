package solution2

import "testing"

func Test_numRollsToTarget(t *testing.T) {
	type args struct {
		n      int
		k      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:      1,
				k:      6,
				target: 3,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				n:      2,
				k:      6,
				target: 7,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				n:      30,
				k:      30,
				target: 500,
			},
			want: 222616187,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRollsToTarget(tt.args.n, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("numRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
