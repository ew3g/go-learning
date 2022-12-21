// Package popcount count the set of bit whose value is 1
package popcount

import "fmt"

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	var value byte
	for i := 0; i < 8; i++ {
		value += pc[byte(x>>(i*8))]
	}
	return int(value)
}

func SecondPopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println(PopCount(10))
	fmt.Println(SecondPopCount(10))
}
