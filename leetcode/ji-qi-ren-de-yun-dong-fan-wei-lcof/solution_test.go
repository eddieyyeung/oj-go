package jiqirendeyundongfanweilcof

import "testing"

func Test_wardrobeFinishing(t *testing.T) {
	type args struct {
		m   int
		n   int
		cnt int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				m:   4,
				n:   7,
				cnt: 5,
			},
			want: 18,
		},
		{
			name: "",
			args: args{
				m:   16,
				n:   8,
				cnt: 4,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wardrobeFinishing(tt.args.m, tt.args.n, tt.args.cnt); got != tt.want {
				t.Errorf("wardrobeFinishing() = %v, want %v", got, tt.want)
			}
		})
	}
}
