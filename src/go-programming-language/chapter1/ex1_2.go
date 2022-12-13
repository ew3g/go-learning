// Print all command-line arguments and its index
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(index) + " " + arg
		sep = " "
	}
	fmt.Println(s)
}
