package solution2

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		x  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s1: "1100011000",
				s2: "0101001010",
				x:  2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.s1, tt.args.s2, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
