// 1039. 多边形三角剖分的最低得分
// https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/description/
package solution

import "testing"

func Test_minScoreTriangulation(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				values: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				values: []int{3, 7, 4, 5},
			},
			want: 144,
		},
		{
			name: "",
			args: args{
				values: []int{1, 3, 1, 4, 1, 5},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minScoreTriangulation(tt.args.values); got != tt.want {
				t.Errorf("minScoreTriangulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
