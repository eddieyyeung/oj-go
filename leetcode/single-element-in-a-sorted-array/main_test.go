package singleelementinasortedarray

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{
					1, 1, 2, 3, 3, 4, 4, 8, 8,
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{3, 3, 7, 7, 10, 11, 11},
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				nums: []int{3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
