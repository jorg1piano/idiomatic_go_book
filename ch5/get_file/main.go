package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")
	f, closer, err := getFile((os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	data := make([]byte, 2048)

	numBytes, err := f.Read(data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bytes", numBytes)
	os.Stdout.Write(data[:numBytes])

}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		fmt.Println("close file")
		file.Close()
	}, nil
}
