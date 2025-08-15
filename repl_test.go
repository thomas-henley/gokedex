package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello  world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: nil,
		},
		{
			input:    "hello\tBEAUTIFUL\n\tworld\n",
			expected: []string{"hello", "beautiful", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("got %v wanted %v", len(actual), len(c.expected))
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("got %v wanted %v", word, expectedWord)
			}
		}
	}
}
