package main

import "fmt"

// with alias
import math2 "golang-book/chapter11/math"

// without alias
// import "golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math2.Average(xs)
	maximum := math2.Maximum(xs)
	minimum := math2.Mininum(xs)
	fmt.Println("Average:", avg)
	fmt.Println("Maximum:", maximum)
	fmt.Println("Minimum:", minimum)
}
