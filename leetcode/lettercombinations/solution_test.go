package lettercombinations_test

import (
	"fmt"
	"oj-go/leetcode/lettercombinations"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.digits), func(t *testing.T) {
			got := lettercombinations.Run(tc.digits)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
