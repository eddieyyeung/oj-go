package solution1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				prices: []int{8, 4, 6, 2, 3},
			},
			want: []int{4, 2, 4, 2, 3},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := finalPrices(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
