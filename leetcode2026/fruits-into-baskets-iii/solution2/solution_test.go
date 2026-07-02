package fruits_into_baskets_iii

import "testing"

func Test_numOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				fruits:  []int{4, 2, 5},
				baskets: []int{3, 5, 4},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				fruits:  []int{3, 6, 1},
				baskets: []int{6, 4, 7},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruits(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
