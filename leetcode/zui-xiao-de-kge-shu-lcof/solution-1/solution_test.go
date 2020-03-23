package solution_1

import (
	"reflect"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{0, 1, 1, 2, 4, 4, 1, 3, 3, 2}, 6}, []int{0, 1, 1, 1, 2, 2}},
		{"", args{[]int{3, 2, 1}, 2}, []int{1, 2}},
		{"", args{[]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8}, []int{0, 0, 1, 1, 2, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
