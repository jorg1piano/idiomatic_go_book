package main

import "fmt"

func main() {
	printSummedValues()
	printOverflowingInts()
}
func printSummedValues() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	fmt.Println(sum3)
	var sum4 byte = byte(x) + b

	fmt.Println(sum3, sum4)
}

func printOverflowingInts() {
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18446744073709551615
	fmt.Println("smallI", smallI+1)
	fmt.Println("bigI", bigI+1)
}
