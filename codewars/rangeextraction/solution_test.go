package rangeextraction

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		list []int
		want string
	}{
		{
			[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20},
			"-6,-3-1,3-5,7-11,14,15,17-20",
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.list), func(t *testing.T) {
			got := Solution(tc.list)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
