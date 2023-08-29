package coursescheduleii

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0}, {1, 2}, {0, 1},
				},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: []int{0, 1},
		},
		{
			name: "",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{3, 2},
				},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
