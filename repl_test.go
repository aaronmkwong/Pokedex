package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input) // Capitalized 'C' to match previous implementation

		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("CleanInput(%q) returned %d words, expected %d", c.input, len(actual), len(c.expected))
			continue // Skip word check if lengths do not match to avoid out-of-bounds panic
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("CleanInput(%q) at index %d: got %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}