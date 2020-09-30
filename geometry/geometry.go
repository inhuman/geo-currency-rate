package geometry

type Point struct {
	X float64
	Y float64
}

// QuadrantDetector detects in which quadrant point placed
// point with coordinates {0, 0} is considered invalid
func QuadrantDetector(p Point) int {

	i := -1

	switch {
	case (p.X > 0) && (p.Y > 0):
		i = 1

	case (p.X < 0) && (p.Y > 0):
		i = 2

	case (p.X < 0) && (p.Y < 0):
		i = 3

	case (p.X > 0) && (p.Y < 0):
		i = 4
	}

	return i
}

// IsInRadius detects that the point placed inside or on a circle
// for this the Pythagorean theorem is used
func IsInRadius(radius float64, point Point) bool {

	if (point.X*point.X)+(point.Y*point.Y) <= radius*radius {
		return true
	}

	return false
}
