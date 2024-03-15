package main

import "fmt"

func main() {
	fmt.Println("Hello pointers")
	dereferenceAndIndirectionOperator()
	fmt.Println()

	nilPointers()
	pointWithNew()
	pointerRepitition()
}

func dereferenceAndIndirectionOperator() {
	x := 10
	pointerX := &x
	valueX := *pointerX
	fmt.Println(pointerX)
	fmt.Println(valueX)
}

func nilPointers() {

	var nilPointer *int

	fmt.Println(nilPointer)

	// Dereferencing a nil pointer will cause a runtime error
	// fmt.Println(*nilPointer)
	// output:
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x1045a8b38]
}

func pointWithNew() {
	var x *int = new(int)
	fmt.Println(*x, "*x should be 0")
	var y = new(int)
	fmt.Println(*y, "*y should be 0")

	var str = new(string)
	fmt.Println(*str, "is an empty string that points to", str)
}

func pointerRepitition() {
	x := 69

	pointerX := &x

	fmt.Println(x, *pointerX, "is the same")
	fmt.Println(&x, pointerX, "is the same")
}

func structPointers() {}
