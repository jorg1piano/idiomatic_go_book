package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if x := 5; x == 5 {
		fmt.Printf("Give me five")
	}

	for i := 0; i < 100; i++ {
		pickRandom()
	}
}

func pickRandom() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Thats to low", n)
	} else if n > 5 {
		fmt.Println("That is to big", n)
	} else {
		fmt.Println("Thats a good number", n)
	}

}

func giveMeFive() int {
	return 5
}
