package androidunlockpatterns

import "testing"

func Test_numberOfPatterns(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"case1",
		// 	args{
		// 		1,
		// 		1,
		// 	},
		// 	9,
		// },
		{
			"case2",
			args{
				1,
				2,
			},
			65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPatterns(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("numberOfPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}
