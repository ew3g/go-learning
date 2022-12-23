// basename2 removes directory components and a .suffix using strings library
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("a"))
	fmt.Println(basename2("a.go"))
	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename2("a/b.c.go"))
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
