package main

import (
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "something failed"
}

// This function returns an error interface.
// Although it returns a `nil` pointer to `MyError`,
// the error interface itself is not `nil` because it has a type (*MyError).
func returnsError() error {
	var err *MyError // Here, err is a nil pointer to MyError.
	return err       // returns an interface with a non-nil type (*MyError) but a nil value.
}

func main() {
	err := returnsError()
	if err != nil {
		fmt.Println("Error:", err) // This will execute because err is not nil.
	} else {
		fmt.Println("No error")
	}
}
