package solution

import (
	"fmt"
	"testing"
)

func Test_fn2(t *testing.T) {
	type args struct {
		left   int
		right  int
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				left:   0,
				right:  7,
				nums:   []int{0, 1, 2, 3, 4, 5, 6, 7},
				target: 6,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			got := fn2(tt.args.left, tt.args.right, tt.args.nums, tt.args.target)
			t.Log("got", got)
		})
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2, 7},
			},
			want: 2.5,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2)
			t.Log("got", got)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
