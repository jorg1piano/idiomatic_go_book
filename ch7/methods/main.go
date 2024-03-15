package main

import "fmt"

func main() {
	p := Person{"Jørgen", "Andersen"}
	fmt.Println("name before reset", p.FullName())
	p.Reset()
	fmt.Println("name after reset", p.FullName())

}

type Person struct {
	firstname string
	lastname  string
}

// When method does not modify the receiver use a value receiver
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstname, p.lastname)
}

// When to use a pointer receiver:
// 1. When the method modifies the receiver (ie person)
// 2. If the method needs to handle nil instances it has to be a receiver
func (p *Person) Reset() {
	p.firstname = ""
	p.lastname = ""
}
