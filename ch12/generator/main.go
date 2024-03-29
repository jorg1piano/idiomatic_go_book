package main

import (
	"context"
	"fmt"
)

// Just an example, dont use go routines and channels like this
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 4 {
			break
		}
		fmt.Println(i)
	}

	cancel()
}

// <-chan int means "return a recieve only channel that produces int's"
func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Done hit")
				return
			case ch <- i:
			}
		}
	}()
	return ch
}
