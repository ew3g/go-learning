package math

import "fmt"

// To check the documentation of this function:
// go doc golang-book/chapter11/math Average

// Average Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Minimum Finds the minimum of series of numbers
func Minimum(xs []float64) float64 {
	minimum := xs[0]
	for _, x := range xs {
		if x < minimum {
			minimum = x
		}
	}
	return minimum
}

// Maximum Finds the maximum of series of numbers
func Maximum(xs []float64) float64 {
	maximum := xs[0]
	for _, x := range xs {
		if x > maximum {
			maximum = x
		}
	}
	return maximum
}

// This function is not visible to external callers of lib math
// because does not start with a capital letter
// If it was NotVisibleExternally(), would be visible externally
func notVisibleExternally() {
	fmt.Println("ok")
}
