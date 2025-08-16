package main

import (
	"testing"
	"time"

	"github.com/thomas-henley/gokedex/internal/pokecache"
)

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

func TestCommandMap(t *testing.T) {
	cache = pokecache.NewCache(5 * time.Second)
	t.Run("get first page of locations", func(t *testing.T) {
		config := Config{}
		expectedLen := 20
		expectedFirst := "canalave-city-area"
		expectedLast := "mt-coronet-1f-from-exterior"

		actual, _ := getNextLocations(&config)

		if len(actual) != expectedLen {
			t.Errorf("expected %v got %v", expectedLen, len(actual))
			return
		}

		if actual[0] != expectedFirst {
			t.Errorf("expected %v got %v", expectedFirst, actual[0])
		}

		if actual[19] != expectedLast {
			t.Errorf("expected %v got %v", expectedLast, actual[19])
		}
	})
}
