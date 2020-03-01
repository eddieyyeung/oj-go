package backwardsreadprimes_test

import (
	"fmt"
	"oj-go/codewars/backwardsreadprimes"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		start int
		stop  int
		want  []int
	}{
		{1, 100, []int{13, 17, 31, 37, 71, 73, 79, 97}},
		{1, 31, []int{13, 17, 31}},
		{501, 599, nil},
		{9900, 10000, []int{9923, 9931, 9941, 9967}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.start, tc.stop), func(t *testing.T) {
			got := backwardsreadprimes.BackwardsPrime(tc.start, tc.stop)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
