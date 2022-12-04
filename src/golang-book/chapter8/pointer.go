package main

import "fmt"

// The zero function takes a pointer of int  a parameter
func zero(xPtr *int) {
	// So the parameter is passed as reference
	//and changes on its memory address affects the external variable
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	// The parameter is given by reference and not by value
	// Changes made by the function affects it here
	fmt.Println(x) // x now is zero
}
