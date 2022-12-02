package main

import "fmt"

func main() {
	// This example program reads an input from the user and doubles it
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
