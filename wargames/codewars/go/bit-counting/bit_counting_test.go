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
		{0, 0},
		{4, 1},
		{7, 3},
		{9, 2},
		{10, 2},
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
