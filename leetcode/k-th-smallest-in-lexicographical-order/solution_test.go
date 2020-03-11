package k_th_smallest_in_lexicographical_order

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "case1", args: args{n: 100, k: 90}, want: 9},
		// {name: "case1", args: args{n: 5210, k: 387}, want: 1346},
		{name: "case2", args: args{n: 13, k: 8}, want: 4},
		// {name: "case3", args: args{n: 10, k: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
