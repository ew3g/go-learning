package main

import "fmt"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	// Shorter statement variable assignment
	y := "Short Statement for creating a variable and assigning it in the same command"
	fmt.Println(y)

	// Constant
	const c string = "This is a constant"
	fmt.Println(c)

	// Defining multiple variables
	var (
		a = 5
		b = 10
		d = 15
	)
	fmt.Println(a, b, d)

	const (
		e = 20
		f = 21
		g = 22
	)
	fmt.Println(e, f, g)
}
