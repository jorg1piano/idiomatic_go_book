package main

import "fmt"

func main() {
	r := []int{1, 2, 0, 3}
	for _, i := range r {
		div60(i)
	}
}
func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
