package solution

import "testing"

func Test_maximumSumOfHeights(t *testing.T) {
	type args struct {
		maxHeights []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				maxHeights: []int{5, 3, 4, 1, 1},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSumOfHeights(tt.args.maxHeights); got != tt.want {
				t.Errorf("maximumSumOfHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
