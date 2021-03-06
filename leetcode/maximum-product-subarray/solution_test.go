package maximumproductsubarray

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{2, 3, -2, 4},
			},
			6,
		},
		{
			"case2",
			args{
				[]int{-2, 0, -1},
			},
			0,
		},
		{
			"[0,2]",
			args{
				[]int{0, 2},
			},
			2,
		},
		{
			"[-2]",
			args{
				[]int{-2},
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
