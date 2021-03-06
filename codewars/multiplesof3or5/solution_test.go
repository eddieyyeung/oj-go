package multiplesof3or5_test

import (
	"fmt"
	"oj-go/codewars/multiplesof3or5"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		number int
		want   int
	}{
		{10, 23},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.number), func(t *testing.T) {
			got := multiplesof3or5.Multiple3And5(tc.number)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
