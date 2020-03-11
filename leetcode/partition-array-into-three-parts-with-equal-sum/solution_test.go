package partition_array_into_three_parts_with_equal_sum

import "testing"

func Test_canThreePartsEqualSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: string("case1"), args: args{A: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}}, want: true},
		{name: string("case2"), args: args{A: []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}}, want: false},
		{name: string("case3"), args: args{A: []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}}, want: true},
		{name: string("case4"), args: args{A: []int{10, -10, 10, -10, 10, -10, 10, -10}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.args.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
