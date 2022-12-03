package main

import "fmt"

func main() {
	fmt.Println(fib(8))
}
func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
