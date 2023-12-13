// 154. 寻找旋转排序数组中的最小值 II
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/description/
package solution

import "testing"

func Test_findMin(t *testing.T) {
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
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
