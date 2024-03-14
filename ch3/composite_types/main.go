package main

import "fmt"

func main() {
	usingMapAsSet()
	emptyStructImpl()
}
func usingMapAsSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 100}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func emptyStructImpl() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = struct{}{}
	}

	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in set")
	}
}
