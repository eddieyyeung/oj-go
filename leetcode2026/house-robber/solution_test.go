package house_robber

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "示例2",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "单个元素",
			nums: []int{5},
			want: 5,
		},
		{
			name: "两个元素",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "全相等",
			nums: []int{3, 3, 3, 3},
			want: 6,
		},
		{
			name: "递增序列",
			nums: []int{1, 2, 3, 4, 5},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
