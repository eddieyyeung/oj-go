package solution2

import "testing"

func Test_maximumInvitations(t *testing.T) {
	type args struct {
		favorite []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "0",
		// 	args: args{
		// 		favorite: []int{2, 2, 1, 2},
		// 	},
		// 	want: 3,
		// },
		{
			name: "",
			args: args{
				favorite: []int{1, 0, 3, 2, 5, 6, 7, 4, 9, 8, 11, 10, 11, 12, 10},
			},
			want: 11,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		favorite: []int{1, 0, 0, 2, 1, 4, 7, 8, 9, 6, 7, 10, 8},
		// 	},
		// 	want: 6,
		// },

		// {
		// 	name: "1",
		// 	args: args{
		// 		favorite: []int{1, 2, 0},
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		favorite: []int{3, 0, 1, 4, 1},
		// 	},
		// 	want: 4,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumInvitations(tt.args.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}
