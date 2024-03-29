package main

import "fmt"

// prints "main: 2, 1"
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGoroutine := 1

		// write to channel 1
		ch1 <- inGoroutine
		// read from ch2
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()

	inMain := 2
	var fromGoroutine int
	select {
	case ch2 <- inMain:
	case fromGoroutine = <-ch1:
	}

	fmt.Println("main:", inMain, fromGoroutine)
}
