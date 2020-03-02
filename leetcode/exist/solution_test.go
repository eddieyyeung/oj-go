package exist_test

import (
	"fmt"
	"oj-go/leetcode/exist"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEE",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCB",
			false,
		},
		{
			[][]byte{
				{'a', 'a'},
			},
			"aaa",
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.board, tc.word), func(t *testing.T) {
			got := exist.Run(tc.board, tc.word)
			if reflect.DeepEqual(got, tc.want) == false {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
