package main

import "fmt"

func main() {
	ex1()
}
func ex1() {
	var pointerCount *Counter
	fmt.Println(pointerCount == nil, "prints true")
	var incrementer Incrementer
	fmt.Println(incrementer == nil, "prints true")
	incrementer = pointerCount
	fmt.Println(incrementer == nil, "prints false")
}

type Incrementer interface{}
type Counter interface{}
