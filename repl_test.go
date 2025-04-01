package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string // Change type from string to []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "   hello     World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if !equal(got, c.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}