package main

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {
	var tests = []struct {
		x    uint
		want int
	}{
		{1234, 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.x), func(t *testing.T) {
			ans := CountBits(test.x)
			if ans != test.want {
				t.Errorf("Correct: %d vs answer: %d", test.want, ans)
			}
		})
	}
}
