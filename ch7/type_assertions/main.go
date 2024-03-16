package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	err := typeAssertion()
	if err != nil {
		fmt.Println(err)
	}
}

// This is not a type conversion. Convertions change a value to a new type.
// While assertions reveal the type of the value stored in the interface.
func typeAssertion() error {
	var i any
	var mine MyInt = 20
	i = mine
	// This is the unsafe way of casting
	// fmt.Println(i.(string))
	// output:
	// panic: interface conversion: interface {} is main.MyInt, not string

	i3, ok := i.(string)
	if !ok {
		return fmt.Errorf("unecpected type for %v", i)
	}
	fmt.Println(i3)
	return nil
}
