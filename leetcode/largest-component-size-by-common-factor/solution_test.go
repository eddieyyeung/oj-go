package largest_component_size_by_common_factor

import (
	"testing"
)

func Test_largestComponentSize(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{66, 100, 37, 40, 41, 76, 49, 62}}, 5},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 6},
		{"", args{[]int{4, 6, 15, 35}}, 4},
		{"", args{[]int{20, 50, 9, 63}}, 2},
		{"", args{[]int{2, 3, 6, 7, 4, 12, 21, 39}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestComponentSize(tt.args.A); got != tt.want {
				t.Errorf("largestComponentSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
