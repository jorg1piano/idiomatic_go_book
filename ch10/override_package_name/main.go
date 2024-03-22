package main

import (
	"fmt"
	"overide_package_name/seed"
)

func main() {
	fmt.Println("Hello from main")

	s := seed.Rand()
	fmt.Println(s)
}
