package main

import "fmt"

// Just an example, dont use go routines and channels like this
func main() {
	for i := range countTo(10) {
		fmt.Println(i)
	}
}

// <-chan int means "return a recieve only channel that produces int's"
func countTo(max int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
