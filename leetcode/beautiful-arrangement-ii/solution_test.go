package beautiful_arrangement_ii

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_constructArray(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				n: 4,
				k: 3,
			},
			want: []int{1, 4, 2, 3},
		},
		{
			args: args{
				n: 3,
				k: 1,
			},
			want: []int{1, 2, 3},
		},
		{
			args: args{
				n: 9,
				k: 8,
			},
			want: []int{1, 9, 2, 8, 3, 7, 4, 6, 5},
		},
		{
			args: args{
				n: 5,
				k: 2,
			},
			want: []int{1, 3, 2, 4, 5},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := constructArray(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
