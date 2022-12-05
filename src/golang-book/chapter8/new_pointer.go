package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	zero(xPtr)
	fmt.Println(xPtr) // x is 0
}
