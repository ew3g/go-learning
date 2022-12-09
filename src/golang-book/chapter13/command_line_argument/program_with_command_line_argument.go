package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// to run
// go run program_with_command_line_argument.go -max=100

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")

	// Parse
	flag.Parse()

	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
