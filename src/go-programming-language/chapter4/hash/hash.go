// Hashes a string to SHA256, SHA384 OR SHA512
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

var shaType = flag.String("shaType", "SHA256", "SHA Type")
var text = flag.String("text", "", "String to hash")

func main() {
	//var text *string
	//var shaType *string

	//s := "teste"
	//t := "SHA512"
	//text = &s
	//shaType = &t
	flag.Parse()
	switch strings.ToUpper(*shaType) {
	case "SHA256":
		fmt.Printf("%x", sha256.Sum256([]byte(*text)))
	case "SHA384":
		fmt.Printf("%x", sha512.Sum384([]byte(*text)))
	case "SHA512":
		fmt.Printf("%x", sha512.Sum512([]byte(*text)))
	default:
		fmt.Printf("%x", sha256.Sum256([]byte(*text)))
	}
}
