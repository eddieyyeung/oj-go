package maximum_total_damage_with_spell_casting

import "testing"

func Test_maximumTotalDamage(t *testing.T) {
	tests := []struct {
		name  string
		power []int
		want  int64
	}{
		{"example1", []int{1, 1, 3, 4}, 6},
		{"example2", []int{7, 1, 6, 6}, 13},
		{"single", []int{5}, 5},
		{"all_same", []int{3, 3, 3}, 9},
		{"no_conflict", []int{1, 4, 7}, 12},
		{"conflict_skip", []int{1, 2, 3}, 3},
		{"", []int{5, 9, 2, 10, 2, 7, 10, 9, 3, 8}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalDamage(tt.power); got != tt.want {
				t.Errorf("maximumTotalDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
