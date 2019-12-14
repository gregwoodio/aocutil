package aocutil

// Coord is a point on a Cartesian plane
type Coord struct {
	X, Y int
}

// GetSlope returns the delta X delta Y values between
// the provided coordinate and the target.
func (source *Coord) GetSlope(c *Coord) (int, int) {
	var dx, dy int

	dx = source.X - c.X
	dy = source.Y - c.Y

	gcd := GreatestCommonDivisor(dx, dy)
	dx /= gcd
	dy /= gcd

	return dx, dy
}
