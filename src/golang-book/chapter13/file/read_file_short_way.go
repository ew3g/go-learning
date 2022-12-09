package main

import (
	"fmt"
	"os"
)

func main() {
	bs, err := os.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
