package main

import "fmt"

func main() {
	panic("PANIC")
	//recover stops the panic state and recover the value given to the panic function
	str := recover()
	fmt.Println(str)
}
