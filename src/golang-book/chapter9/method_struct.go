package main

import (
	"fmt"
	"math"
)

type Circle2 struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Method
// a function which between func and its name we add the receiver
// the parameter c of type *Circle2
// We will be able to call area directly from object c, like c.area()
// With this, we dont need reference the object as pointer parameter
// Go does this automatically, so any changes inside the objects, affects it outside
func (c *Circle2) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	c := Circle2{1, 2, 3}
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
