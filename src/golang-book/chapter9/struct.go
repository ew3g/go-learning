package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// struct initialization
func main() {
	//empty circle
	var c1 Circle
	fmt.Println(c1)
	// new empty pointer of Circle
	c2 := new(Circle)
	fmt.Println(c2)
	// with named values
	c3 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c3)
	// without names
	c4 := Circle{0, 0, 5}
	fmt.Println(c4)

	// Accessing fields
	fmt.Println(c4.x, c4.y, c4.r)
	c4.x = 10
	c4.y = 5

	fmt.Println(circleArea(c4))
	fmt.Println(circleAreaPointer(&c4))
}

// Copying parameters
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// Referencing parameters with pointers
func circleAreaPointer(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// Method
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
