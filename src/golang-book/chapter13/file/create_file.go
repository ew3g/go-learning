package main

import "os"

func main() {
	file, err := os.Create("test2.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("test create file")
}
