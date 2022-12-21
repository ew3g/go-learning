// Package lengthconv performs Feet and Meter conversion
package lengthconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g feets", f) }
func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }
