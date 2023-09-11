package coursescheduleiii

import "testing"

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				courses: [][]int{
					{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {10, 19}, {4, 18},
				},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				courses: [][]int{
					{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19},
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				courses: [][]int{
					{5, 5},
					{4, 6},
					{2, 6},
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				courses: [][]int{
					{100, 200},
					{200, 1300},
					{1000, 1250},
					{2000, 3200},
				},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				courses: [][]int{
					{1, 2},
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				courses: [][]int{
					{3, 2},
					{4, 3},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scheduleCourse(tt.args.courses); got != tt.want {
				t.Errorf("scheduleCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}
