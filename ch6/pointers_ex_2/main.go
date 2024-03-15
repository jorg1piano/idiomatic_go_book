package main

import "fmt"

func main() {
	names := []string{"Jørgen", "Jonas", "marlene"}
	UpdateSlice(names, "Zero")
	fmt.Println(names, "outside UpdateSlice")

	GrowSlice(names, "marlene")
	fmt.Println(names)
	fmt.Println(names, "outside GrowSlice. Changes are not visible")

	GrowSliceAndMutate(names, "test1", "test2")

	fmt.Println(names, "outside GrowSliceAndMutate")
}

func UpdateSlice(strs []string, value string) {
	strs[len(strs)-1] = value
	fmt.Println(cap(strs), "is 3")
	fmt.Println(strs, "inside UpdateSlice")
}

func GrowSlice(strs []string, value string) {
	strs = append(strs, value)
	fmt.Println(strs, "inside GrowSlice")
}

func GrowSliceAndMutate(strs []string, val1 string, val2 string) {
	// We are getting a new slice here with a new capacity that is double the old one
	// Since the values are copied over, change the strs[0] to a new string
	// does not change the original passed in slice
	fmt.Println(cap(strs), "capacity before")
	strs = append(strs, val1)
	fmt.Println(cap(strs), "capacity after")
	strs[0] = val1
}
