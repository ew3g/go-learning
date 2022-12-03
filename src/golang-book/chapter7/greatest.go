package main

import "fmt"

func greatest(args ...float64) float64 {
	greatestValue := args[0]
	for _, value := range args {
		if value > greatestValue {
			greatestValue = value
		}
	}
	return greatestValue
}

func main() {
	fmt.Println(greatest(1, 2, 3, 4))
	// ... unpack the slice into variables
	fmt.Println(greatest([]float64{1, 2, 3, 4}...))
}
