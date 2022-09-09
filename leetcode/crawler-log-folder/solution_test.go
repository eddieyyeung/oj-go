package crawler_log_folder

import (
	"fmt"
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		args    args
		wantAns int
	}{
		{
			args: args{
				logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			},
			wantAns: 2,
		},
		{
			args: args{
				logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			},
			wantAns: 3,
		},
		{
			args: args{
				logs: []string{"d1/", "../", "../", "../"},
			},
			wantAns: 0,
		},
		{
			args: args{
				logs: []string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"},
			},
			wantAns: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if gotAns := minOperations(tt.args.logs); gotAns != tt.wantAns {
				t.Errorf("minOperations() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
