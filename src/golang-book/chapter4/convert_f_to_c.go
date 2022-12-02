package main

import "fmt"

func main() {
	// Program that takes a Fahrenheit temperature and converts it to Celsius
	fmt.Println("Enter the temperature in ºF:")

	var fahrenheitTemperature float64
	fmt.Scanf("%f", &fahrenheitTemperature)

	celsiusTemperature := (fahrenheitTemperature - 32) * 5 / 9
	fmt.Println("The temperature in Celsius is: ", celsiusTemperature)
}
