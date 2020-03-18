package simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"/home/"}, "/home"},
		{"", args{"/../"}, "/"},
		{"", args{"/home//foo/"}, "/home/foo"},
		{"", args{"/a/./b/../../c/"}, "/c"},
		{"", args{"/a/../../b/../c//.//"}, "/c"},
		{"", args{"/a//b////c/d//././/.."}, "/a/b/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
