package ioutil

import (
	"io"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}