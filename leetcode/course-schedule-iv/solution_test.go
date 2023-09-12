package coursescheduleiv

import (
	"reflect"
	"testing"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "",
			args: args{
				numCourses:    5,
				prerequisites: [][]int{{4, 3}, {4, 1}, {4, 0}, {3, 2}, {3, 1}, {3, 0}, {2, 1}, {2, 0}, {1, 0}},
				queries:       [][]int{{1, 4}, {4, 2}, {0, 1}, {4, 0}, {0, 2}, {1, 3}, {0, 1}},
			},
			want: []bool{false, true, false, true, false, false, false},
		},
		{
			name: "",
			args: args{
				numCourses:    5,
				prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
				queries:       [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}},
			},
			want: []bool{true, false, true, false},
		},
		{
			name: "",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
				queries:       [][]int{{0, 1}, {1, 0}},
			},
			want: []bool{false, true},
		},
		{
			name: "",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{},
				queries:       [][]int{{1, 0}, {0, 1}},
			},
			want: []bool{false, false},
		},
		{
			name: "",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
				queries:       [][]int{{1, 0}, {1, 2}},
			},
			want: []bool{true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrerequisite(tt.args.numCourses, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkIfPrerequisite() = %v, want %v", got, tt.want)
			}
		})
	}
}
