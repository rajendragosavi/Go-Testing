package main

import (
	"testing"
)

func TestCalculatesum(t *testing.T) {

	tests := []struct {
		x int
		y int
		n int
	}{
		{1, 2, 3},
		{4, 6, 10},
		{15, 23, 30},
		{1, -1, 0},
	}

	for _, test := range tests {
		total := Calculatesum(test.x, test.y)
		if total != test.n {
			t.Errorf("Error occured, got: %d , want: %d", test.n, total)
		}
	}
}
