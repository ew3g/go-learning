package main

import "fmt"

func main() {
	// This function is executed with defer
	// So in the end of execution of main()
	// it will be executed
	// after the panic and capture the value from panic with defer
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
