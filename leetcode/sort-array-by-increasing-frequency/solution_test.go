package sort_array_by_increasing_frequency

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{1, 1, 2, 2, 2, 3},
			},
			want: []int{3, 1, 1, 2, 2, 2},
		},
		{
			args: args{
				nums: []int{2, 3, 1, 3, 2},
			},
			want: []int{1, 3, 3, 2, 2},
		},
		{
			args: args{
				nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			},
			want: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := frequencySort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
