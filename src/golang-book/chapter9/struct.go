package main

type Circle struct {
	x, y, r float64
}

// struct initialization
func main() {
	//empty circle
	var c1 Circle
	// new empty pointer of Circle
	c2 := new(Circle)
	// with named values
	c3 := Circle{x: 0, y: 0, r: 5}
	// without names
	c4 := Circle{0, 0, 5}

	// Access
	
}