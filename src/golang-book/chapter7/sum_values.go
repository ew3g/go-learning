package main

import "fmt"

func main() {
	fmt.Println(sum([]float64{1, 2, 3, 4, 5}))
}

func sum(x []float64) float64 {
	total := 0.0
	for _, value := range x {
		total += value
	}
	return total
}
