package solution2

import "testing"

func Test_countTexts(t *testing.T) {
	type args struct {
		pressedKeys string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				pressedKeys: "22233",
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				pressedKeys: "344644885",
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				pressedKeys: "222222222222222222222222222222222222",
			},
			want: 82876089,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTexts(tt.args.pressedKeys); got != tt.want {
				t.Errorf("countTexts() = %v, want %v", got, tt.want)
			}
		})
	}
}
