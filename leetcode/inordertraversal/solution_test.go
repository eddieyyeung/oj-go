package inordertraversal

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					8,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				6,
				nil,
				nil,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{
				root,
			},
			[]int{4, 2, 7, 5, 8, 1, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
