package longestcommonsubsequence_test

import (
	"fmt"
	"oj-go/codewars/longestcommonsubsequence"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		s1   string
		s2   string
		want string
	}{
		{"a", "b", ""},
		{"abcdef", "abc", "abc"},
		{"132535365", "123456789", "12356"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.s1, tc.s2), func(t *testing.T) {
			got := longestcommonsubsequence.LCS(tc.s1, tc.s2)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
