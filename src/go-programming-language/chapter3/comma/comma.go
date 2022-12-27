// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(comma("123"))
	fmt.Println(commaWithBuffer("123"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithBuffer(s string) string {
	var buf bytes.Buffer
	for _, v := range s {
		buf.WriteRune(v)
		if len(buf.String()) == 3 && len(s) > 3 && !strings.Contains(buf.String(), ",") {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
