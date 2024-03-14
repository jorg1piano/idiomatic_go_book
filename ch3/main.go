package main

import "fmt"

func main() {
	example1()
	example2()
	example3SharedMemory()
	example4()
	example5()
	example6CompositeTypes()
	example7SliceToArray()
	exampleArrayPointers()
}

func example1() {
	fmt.Println("example1")
	var x []int
	fmt.Println(x)
	fmt.Println(len(x))

	// fmt.Println(x[0])
	// output
	// panic: runtime error: index out of range [0] with length 0
}

func example2() {
	x := []int{1, 2, 3}
	fmt.Println(x)

	// fmt.Println(x[4])
	// output
	// panic: runtime error: index out of range [4] with length 3
}

func example3SharedMemory() {
	fmt.Println("Example 3 - Create a slice out of the previous slice that points to the same memory")
	x := []int{2, 4, 6, 8}
	c := x[1:3]
	fmt.Println("Copy using [:]")

	x[1] = 69
	x[2] = 69
	fmt.Println(x)
	fmt.Println(c)
}

func example4() {
	fmt.Println("Example 4 - Create a copy")
	src := []string{"hello", "hi", "goodby"}
	dst := make([]string, len(src))
	copy(dst, src)
	dst[2] = "goodbye"

	fmt.Println(src)
	fmt.Println(dst)
}

func example5() {
	a := make([]int, 5)

	fmt.Println("Prints the zero values")
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
	fmt.Println(a[4])

	// fmt.Println(a[5])
	// output
	// panic: runtime error: index out of range [5] with length 5
}

func example6CompositeTypes() {
	fmt.Println("example6CompositeTypes")
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)

	copy(y, d[:])

	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
}

func example7SliceToArray() {
	// When converting a slice to an array the slice is copied into new memory.
	// That means that changing slice wont affect the array, and vice versa
	fmt.Println("example7SliceToArray")
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	xArray[0] = 10

	fmt.Println("xSlice")
	fmt.Println(xSlice)
	fmt.Println("xArray")
	fmt.Println(xArray)

	// If the array size is bigger than the slice, the program will panic
	// panicArray := [5]int(xSlice)
	// fmt.Println(panicArray)
	// output
	// panic: runtime error: cannot convert slice with length 4 to array or pointer to array with length 5

}

func exampleArrayPointers() {
	fmt.Println("Array pointers")
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	xArrayPointer[1] = 10
	xSlice[0] = 20

	fmt.Println(xArrayPointer)
	fmt.Println(xSlice)
}
