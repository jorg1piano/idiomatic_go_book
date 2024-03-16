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
	switch j := i.(type) {
	case nil:
		fmt.Println(j, "The value was nil")
	case string:
		fmt.Println(j, "This was a string")
	case int:
		fmt.Println(j, "This was a int")
	default:
		fmt.Println(j, "Cant handle this type here")
	}
}
