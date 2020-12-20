package aocutil

import "testing"

func TestReverse(t *testing.T) {
	testData := []struct {
		input    string
		expected string
	}{
		{
			input:    "abcdef",
			expected: "fedcba",
		},
		{
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			input:    "amanaplanacanalpanama",
			expected: "amanaplanacanalpanama",
		},
	}

	for _, td := range testData {
		actual := reverse(td.input)
		if actual != td.expected {
			t.Errorf("expected '%s' but was '%s'", td.expected, actual)
		}
	}
}
