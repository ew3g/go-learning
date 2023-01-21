package main

import "fmt"

type Celsius float64

func (c *Celsius) String() string {
	return fmt.Sprintf("%.2f ºC", *c)
}

// * celsiusFlag satisfies the flag.Value interface
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "ºC":
		f.Celsius = Celsius(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func main() {
	var c Celsius = 1
	fmt.Println(c.String())
}
