package main

import "fmt"

func main() {
	doThings(nil)
	doThings("mystring")
	doThings(23)

	i := 0
	doThings(&i)

}

func doThings(i any) {
	switch i := i.(type) {
	case nil:
		fmt.Println(i, "The value was nil")
	case string:
		fmt.Println(i, "This was a string")
	case int:
		fmt.Println(i, "This was a int")
	default:
		fmt.Println(i, "Cant handle this type here")
	}
}
