package geometry

import "math"

// Point represents x and y
type Point struct {
	X, Y float64
}

// Distance gets the distance between 2 points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance gets the distance between 2 points - Distance is a method of type Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
