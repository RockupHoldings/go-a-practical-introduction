package surface

import "math"

// Circle calculates and returns the surface area of a circle
func Circle(r float64) float64 {
	return math.Pi * r * r
}
