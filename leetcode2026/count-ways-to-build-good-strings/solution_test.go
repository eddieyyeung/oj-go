package count_ways_to_build_good_strings

import "testing"

func Test_countGoodStrings(t *testing.T) {
	type args struct {
		low  int
		high int
		zero int
		one  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				low:  3,
				high: 3,
				zero: 1,
				one:  1,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				low:  2,
				high: 3,
				zero: 1,
				one:  2,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				low:  200,
				high: 200,
				zero: 10,
				one:  1,
			},
			want: 764262396,
		},
		{
			name: "",
			args: args{
				low:  500,
				high: 500,
				zero: 5,
				one:  2,
			},
			want: 873327137,
		},
		{
			name: "",
			args: args{
				low:  1,
				high: 100000,
				zero: 1,
				one:  1,
			},
			want: 215447031,
		},
		{
			name: "",
			args: args{
				low:  50000,
				high: 100000,
				zero: 50000,
				one:  50000,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				low:  50000,
				high: 100000,
				zero: 2,
				one:  3,
			},
			want: 797774039,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodStrings(tt.args.low, tt.args.high, tt.args.zero, tt.args.one); got != tt.want {
				t.Errorf("countGoodStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
