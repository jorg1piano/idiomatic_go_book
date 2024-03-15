package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	totalBytes, err := readFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalBytes)
}

func readFile(filename string) (totalBytes int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buffer := make([]byte, 4)
	for {
		count, err := file.Read(buffer)

		fmt.Print(string(buffer[:count]))
		if err != nil {
			if err == io.EOF {
				return totalBytes, nil
			}
		}
		totalBytes += count
	}
}
