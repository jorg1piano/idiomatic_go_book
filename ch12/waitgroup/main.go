package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		fmt.Println("First group")
		defer wg.Done()
	}()
	go func() {
		fmt.Println("second group")
		defer wg.Done()
	}()
	go func() {
		fmt.Println("third group")
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Println("All wait groups finished")
}
