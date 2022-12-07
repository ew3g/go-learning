package main

import (
	"fmt"
	"math"
)

type Circle3 struct {
	x, y, r float64
}

type Rectangle3 struct {
	x1, y1, x2, y2 float64
}

func (c *Circle3) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle3) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func distance3(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle3) area() float64 {
	l := distance3(r.x1, r.y1, r.x1, r.y2)
	w := distance3(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle3) perimeter() float64 {
	l := distance3(r.x1, r.y1, r.x1, r.y2)
	w := distance3(r.x1, r.y1, r.x2, r.y1)

	return (l * w) * (l * w)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	circle := Circle3{1, 2, 3}
	rectangle := Rectangle3{1, 2, 3, 4}
	//shapes := []Shape{&circle, &rectangle}

	fmt.Println("circle perimeter: ", circle.perimeter())
	fmt.Println("rectangle perimeter", rectangle.perimeter())
	//fmt.Println("total area:", totalArea(shapes...))
	fmt.Println("total area:", totalArea(&circle, &rectangle))

}
