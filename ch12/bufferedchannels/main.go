package main

import "fmt"

func processChannel(ch chan int) []int {
	const conc = 10
	// make a buffered channels with the size of 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			// write from channel into v
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out

}

func process(i int) int {
	return i * 2
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
	}()

	result := processChannel(ch)
	fmt.Println(result, len(result))
}
