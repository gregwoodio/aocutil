package aocutil

// Coord is a point on a Cartesian plane
type Coord struct {
	X, Y int
}

// GetSlope returns the delta X delta Y values between
// the provided coordinate and the target.
func (source *Coord) GetSlope(c *Coord) (int, int) {
	var dx, dy int

	dx = c.X - source.X
	dy = c.Y - source.Y

	if dx == 0 && dy == 0 {
		return 0, 0
	} else if dx == 0 {
		return 0, 1
	} else if dy == 0 {
		return 1, 0
	}

	gcd := GreatestCommonDivisor(dx, dy)
	dx /= gcd
	dy /= gcd

	return dx, dy
}
