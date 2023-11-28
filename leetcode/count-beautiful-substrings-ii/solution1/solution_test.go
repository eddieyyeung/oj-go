package solution

import "testing"

func Test_beautifulSubstrings(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				s: "uzuxpzou",
				k: 3,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "baeyh",
				k: 2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "abba",
				k: 1,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "bcdf",
				k: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulSubstrings(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("beautifulSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
