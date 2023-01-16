package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// A Path is journey connection the points with straight lines
type Path []Point

// Distance returns the distance traveled along the path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Usage
//perim := geometry.Path{{1, 1}, {5,1}, {5, 4}, {1, 1}}
//fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
//fmt.Println(geometry.Distance()) // "12", method of geometry.Path
