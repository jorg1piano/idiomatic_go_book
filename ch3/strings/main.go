package main

import "fmt"

func main() {
	accessIndex()
	sliceExpressionStrings()
}

func accessIndex() {
	var s string = "Hello there"
	var b byte = s[6]

	fmt.Println(s)
	fmt.Println(b)
}

func sliceExpressionStrings() {
	var name string = "Jorgen Andersen"
	var first string = name[:6]
	var last string = name[8:]
	var firstLetter = name[0:1]
	var fl12 = name[1:2]

	var sunEmoji = "Hello ☀️"

	fmt.Println(first)
	fmt.Println(last)
	fmt.Println(firstLetter)
	fmt.Println(fl12)
	fmt.Println(sunEmoji)
	fmt.Println(len(sunEmoji))
}
