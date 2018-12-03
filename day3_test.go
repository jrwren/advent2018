package main

import (
	"testing"
)

func TestDay3(t *testing.T) {
	t.Run("given part1", func(t *testing.T) {
		input := `#1 @ 1,3: 4x4
		#2 @ 3,1: 4x4
		#3 @ 5,5: 2x2`
		overlapping := day3(input)
		if overlapping == 0 {
			t.Errorf("got %d expected ?", overlapping)
		}
	})
}

func day3(input string) int {
	return 0
}
