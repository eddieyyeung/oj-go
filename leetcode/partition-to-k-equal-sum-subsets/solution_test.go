package partition_to_k_equal_sum_subsets

import "testing"

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{2, 2, 2, 2, 3, 4, 5}, 4}, false},
		{"", args{[]int{4, 3, 2, 3, 5, 2, 1}, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
