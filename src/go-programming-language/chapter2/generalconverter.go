// General converter, get input from parameter or std input
package main

import (
	"bufio"
	"fmt"
	"go-programming-language/chapter2/lengthconv"
	"go-programming-language/chapter2/tempconv"
	"go-programming-language/chapter2/weightconv"
	"os"
	"strconv"
)

func main() {
	values := os.Args[1:]
	if len(values) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			convert(input.Text())
		}
	} else {
		for _, arg := range values {
			convert(arg)
		}
	}
}

func convert(arg string) {
	value, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on parsing: %v\n", err)
	}
	f := tempconv.Fahrenheit(value)
	c := tempconv.Celsius(value)
	meter := lengthconv.Meter(value)
	feet := lengthconv.Feet(value)
	kilogram := weightconv.Kilogram(value)
	pound := weightconv.Pound(value)
	fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", meter, lengthconv.MToF(meter), feet, lengthconv.FToM(feet))
	fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", kilogram, weightconv.KToL(kilogram), pound, weightconv.LToK(pound))
}
