package main

import (
	"bytes"
	"testing"
)

var input_001 string = `
12 3
1 0 1 0 1 1 1 1 0 0 0 0
6 7
12 1
9 7
1 4
10 7
7 1
11 8
5 1
3 7
5 8
4 2
`

var input_002 string = `
4 1
1 1 0 0
1 2
1 3
1 4
`

func Test_solve(t *testing.T) {
	type args struct {
		ca CaseArg
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				ca: scanCase(bytes.NewBufferString(input_002)),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				ca: scanCase(bytes.NewBufferString(input_001)),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.ca); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
