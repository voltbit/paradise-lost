package main

import "testing"

func TestDetectCapital(t *testing.T) {
	var tests = []struct {
		in       string
		expected bool
	}{
		{"HELLO", true},
		{"hellO", false},
		{"HELLo", false},
		{"Hello", true},
		{"mL", false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := DetectCapital(test.in)
			if res != test.expected {
				t.Errorf("got %t for %s | expected %t", res, test.in, test.expected)
			}
		})
	}
}
