package main

import "fmt"

func main() {
	//Integer
	fmt.Println("1 + 1 = ", 1+1)
	//Floating Point
	fmt.Println("1 + 1 = ", 1.0+1.0)

	//String
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[3])
	fmt.Println("Hello " + "World")

	//Boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Exercise
	// Multiplying numbers
	fmt.Println("The result of 21132 x 42452 = ", 32132*42452)

	// What is the output?
	// true
	fmt.Println((true && false) || (false && true) || !(false && false))
}
