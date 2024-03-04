package solution

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, -2, 3, -2},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{5, -3, 5},
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				nums: []int{3, -2, 2, -3},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{-3, -2, -3},
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
