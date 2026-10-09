package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected  []string
	}{
		{
			input: "  hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello, World  ",
			expected: []string{"hello,", "world"},
		},
	
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Return slice not equal to expected slice\nActual: %v(%v)\nExpected: %v(%v)", actual, len(actual), c.expected, len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			
			if word != expected {
				t.Errorf("input: %s\nActual: %v\nExpected: %v", c.input, word, expected)
			}
		}


	}
}
