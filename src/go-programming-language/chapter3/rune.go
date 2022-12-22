// List all the runes in unicode string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "\tHello, BF\t"
	fmt.Println(s)
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
