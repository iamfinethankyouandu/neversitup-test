package main

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "a",
			want:  []string{"a"},
		},
		{
			input: "ab",
			want:  []string{"ab", "ba"},
		},
		{
			input: "abc",
			want:  []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			input: "aabb",
			want:  []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, tt := range tests {
		t.Run("Permutations test", func(t *testing.T) {
			if got := Permutations(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
