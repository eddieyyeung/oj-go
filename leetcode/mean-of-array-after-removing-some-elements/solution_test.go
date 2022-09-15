package mean_of_array_after_removing_some_elements

import (
	"fmt"
	"testing"
)

func Test_trimMean(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args: args{
				arr: []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			},
			want: 2,
		},
		{
			args: args{
				arr: []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0},
			},
			want: 4,
		},
		{
			args: args{
				arr: []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4},
			},
			want: 4.77778,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := trimMean(tt.args.arr); got != tt.want {
				t.Errorf("trimMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
