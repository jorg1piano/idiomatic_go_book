package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGorutine := 1
		ch1 <- inGorutine
		fromMain := <-ch2
		fmt.Println("goroutine:", inGorutine, fromMain)
	}()

	inMain := 2
	ch2 <- inMain
	fromGoroutine := <-ch1
	fmt.Println("main:", inMain, fromGoroutine)
}
