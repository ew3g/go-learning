// Dup2 prints the count, the text of lines and the filename that appear more than once
// in the input. It reads from stdin or from a list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data struct {
	count int
	file  string
}

func main() {
	counts := make(map[string]Data)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, n.file)
		}
	}
}

func countLines(f *os.File, counts map[string]Data) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		data := counts[input.Text()]
		data.file = f.Name()
		data.count++
		counts[input.Text()] = data
	}
	//NOTE: ignoring potential errors from input.Err()
}
