package main

import "fmt"

func main() {
	fmt.Println(half(0))
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(x int64) (value int64, even bool) {
	value = x
	even = x%2 == 0
	return
}
