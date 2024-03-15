package main

import "fmt"

type Person struct {
	firstName, lastName string
	age                 int
}

// compile with go build -gcflags="-m"
func main() {

	p1 := makePerson("jørgen", "andersen", 31)
	// p1 suprisingly escapes to the heap because
	// Println take in a interface{}.
	// Interface values are stored on the heap
	fmt.Println(p1)
	p2 := MakePersonPointer("Jørgen", "Andersen", 31)
	fmt.Println(*p2)
}

func makePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func MakePersonPointer(firstname, lastname string, age int) *Person {
	return &Person{
		firstName: firstname,
		lastName:  lastname,
		age:       age,
	}

}
