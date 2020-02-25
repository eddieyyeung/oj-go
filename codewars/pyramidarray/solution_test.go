package pyramidarray_test

import (
	"fmt"
	"oj-go/codewars/pyramidarray"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want [][]int
	}{
		{0, [][]int{}},
		{1, [][]int{[]int{1}}},
		{2, [][]int{[]int{1}, []int{1, 1}}},
		{3, [][]int{[]int{1}, []int{1, 1}, []int{1, 1, 1}}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n: %v", tc.n), func(t *testing.T) {
			got := pyramidarray.Pyramid(tc.n)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %s; want %s", fmt.Sprint(got), fmt.Sprint(tc.want))
			}
		})
	}
}
