package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// This cannot proceed until ch1 is read and the main goroutine cannot proceed until ch2 is read
	go func() {
		inGorutine := 1
		// Write inGorutine to ch1
		ch1 <- inGorutine
		// read variable from ch2 into fromMain
		fromMain := <-ch2
		fmt.Println("goroutine:", inGorutine, fromMain)
	}()

	inMain := 2
	// write inMain to ch2
	ch2 <- inMain

	// We never get here (read from ch1)
	fromGoroutine := <-ch1
	fmt.Println("main:", inMain, fromGoroutine)
}
