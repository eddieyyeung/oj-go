package countsubstringsthatsatisfykconstraintii

import (
	"reflect"
	"testing"
)

func Test_countKConstraintSubstrings(t *testing.T) {
	type args struct {
		s       string
		k       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "",
			args: args{
				s: "001110001",
				k: 3,
				queries: [][]int{
					{1, 8},
				},
			},
			want: []int64{35},
		},
		{
			name: "",
			args: args{
				s: "0001111",
				k: 2,
				queries: [][]int{
					{0, 6},
				},
			},
			want: []int64{26},
		},
		{
			name: "",
			args: args{
				s:       "010101",
				k:       1,
				queries: [][]int{{0, 5}, {1, 4}, {2, 3}},
			},
			want: []int64{15, 9, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKConstraintSubstrings(tt.args.s, tt.args.k, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countKConstraintSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
