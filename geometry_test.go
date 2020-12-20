package aocutil

import "testing"

func TestGetSlope(t *testing.T) {
	a := Coord{5, 9}
	b := Coord{13, 11}

	dx, dy := b.GetSlope(&a)
	if dx != 4 && dy != 1 {
		t.Errorf("Expected delta X to be %d and delta Y to be %d, but were %d and %d\n", 4, 1, dx, dy)
	}
}
