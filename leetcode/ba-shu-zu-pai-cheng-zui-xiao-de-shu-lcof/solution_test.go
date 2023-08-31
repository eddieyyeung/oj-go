package bashuzupaichengzuixiaodeshulcof

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				nums: []int{10, 2, 2147483647, 214748364, 2147483646},
			},
			want: "10214748364214748364621474836472",
		},
		{
			name: "",
			args: args{
				nums: []int{10, 2},
			},
			want: "102",
		},
		{
			name: "",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "3033459",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
