package main

import "testing"

func TestSanitizeInput(t *testing.T) {
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
			input: "HeLlO World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, test := range cases {
		actual := sanitizeInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("ERROR: lengths of %v and %v are not equal", actual, test.expected)
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := test.expected[i]
			if actualWord != expectedWord {
				t.Errorf("ERROR: %v is not the same as %v", actualWord, expectedWord)
			}
		}
	}
}
