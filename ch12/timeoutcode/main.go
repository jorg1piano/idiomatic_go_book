package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	exampleSuccess()
	exampleTimeout()
}

func exampleSuccess() {
	fmt.Println("Hello")

	result, err := timeLimit[string](timeOutFunction, time.Second*2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got result", result)
}
func exampleTimeout() {
	fmt.Println("Hello")

	result, err := timeLimit[string](timeOutFunction, time.Millisecond*500)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got result", result)
}

func timeOutFunction() string {
	time.Sleep(time.Duration(time.Second * 1))
	return "this will not be recived"
}

func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T, 1)
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()
	go func() {
		out <- worker()
	}()

	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		var zero T
		return zero, errors.New("work timed out")
	}
}
