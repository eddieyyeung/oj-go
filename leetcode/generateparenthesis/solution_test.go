package generateparenthesis_test

import (
	"fmt"
	"oj-go/leetcode/generateparenthesis"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want []string
	}{
		{3, []string{"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.n), func(t *testing.T) {
			got := generateparenthesis.Run(tc.n)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
