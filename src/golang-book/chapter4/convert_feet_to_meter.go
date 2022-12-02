package main

import "fmt"

func main() {
	// Program that takes a value in feets and converts into meters
	fmt.Println("Enter the value in feets:")

	var feetsValue float64
	fmt.Scanf("%f", &feetsValue)

	metersValue := feetsValue * 0.3048
	fmt.Println("The value in meters is: ", metersValue)
}
