package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected %v, got %v", len(c.expected), len(actual))
		}
		for i, actualWord := range actual {
			if actualWord != c.expected[i] {
				t.Errorf("expected %v, got %v", c.expected[i], actualWord)
			}
		}
	}
}
