package count_number_of_texts

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTexts(tt.args.pressedKeys); got != tt.want {
				t.Errorf("countTexts() = %v, want %v", got, tt.want)
			}
		})
	}
}
