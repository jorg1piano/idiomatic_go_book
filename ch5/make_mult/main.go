package main

import "fmt"

func main() {
	oneBase := makeMult(1)
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 5; i++ {
		fmt.Println(oneBase(i), twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
