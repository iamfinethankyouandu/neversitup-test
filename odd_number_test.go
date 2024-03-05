package main

import "testing"

func TestFindOddNumber(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{7},
			want:  7,
		},
		{
			input: []int{0},
			want:  0,
		},
		{
			input: []int{1, 1, 2},
			want:  2,
		},
		{
			input: []int{0, 1, 0, 1, 0},
			want:  0,
		},
		{
			input: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run("odd number", func(t *testing.T) {
			if got := FindOddNumber(tt.input); got != tt.want {
				t.Errorf("FindOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
