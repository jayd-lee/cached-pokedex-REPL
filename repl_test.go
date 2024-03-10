package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The lengths are not euqal: %v vs %v",
				len(actual),
				len(c.expected))
			continue
		}
		// Safe indexing both slices since same length is confirmed
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}
}
