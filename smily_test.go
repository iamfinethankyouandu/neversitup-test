package main

import "testing"

func TestCountSmilyFace(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{":)", ";(", ";}", ":-D"},
			want:  2,
		},
		{
			input: []string{";D", ":-(", ":-)", ";~)"},
			want:  3,
		},
		{
			input: []string{";]", ":[", ";*", ":$", ";-D"},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run("CountSmilyFac", func(t *testing.T) {
			if got := CountSmilyFace(tt.input); got != tt.want {
				t.Errorf("CountSmilyFace() = %v, want %v", got, tt.want)
			}
		})
	}
}
