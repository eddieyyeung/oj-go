package combination_sum_iii

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				k: 3,
				n: 7,
			},
			want: [][]int{{4, 2, 1}},
		},
		{
			name: "",
			args: args{
				k: 3,
				n: 9,
			},
			want: [][]int{
				{4, 3, 2},
				{5, 3, 1},
				{6, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
