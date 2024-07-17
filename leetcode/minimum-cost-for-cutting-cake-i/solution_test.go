package minimumcostforcuttingcakei

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				m:             3,
				n:             2,
				horizontalCut: []int{1, 3},
				verticalCut:   []int{5},
			},
			want: 13,
		},
		{
			name: "",
			args: args{
				m:             2,
				n:             2,
				horizontalCut: []int{7},
				verticalCut:   []int{4},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.m, tt.args.n, tt.args.horizontalCut, tt.args.verticalCut); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
