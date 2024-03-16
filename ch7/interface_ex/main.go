package main

import (
	"compress/gzip"
	"os"
)

func main() {
	readFile()
}

func readFile() (*gzip.Reader, error) {
	fileName := os.Args[1]

	r, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	// r is a *os.File which implements the io.Reader interface
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	return gz, nil
}
