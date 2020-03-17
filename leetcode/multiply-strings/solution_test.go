package multiply_strings

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"123", "456"}, "56088"},
		{"", args{"2", "3"}, "6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"123", "79"}, "202"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mulOne(t *testing.T) {
	type args struct {
		one int
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{9, "98387"}, "885483"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mulOne(tt.args.one, tt.args.num); got != tt.want {
				t.Errorf("mulOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
