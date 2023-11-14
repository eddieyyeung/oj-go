package knapsack

import "testing"

func Test_zero_one_knapsack(t *testing.T) {
	type args struct {
		capacity int
		w        []int
		v        []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				capacity: 4,
				w:        []int{4, 5, 1},
				v:        []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				capacity: 3,
				w:        []int{4, 5, 6},
				v:        []int{1, 2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zero_one_knapsack(tt.args.capacity, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("zero_one_knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
