package main

import "fmt"

func main() {
	f1 := prefixer("Mr cool")
	f2 := prefixer("Mr ass")

	v1 := f1("Jonas")
	v2 := f2("Jorgen")

	fmt.Println(v1)
	fmt.Println(v2)
}

func prefixer(prefix string) func(string) string {
	return func(name string) string {
		return prefix + " " + name
	}
}
