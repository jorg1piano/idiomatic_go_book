package main

type Person struct {
	name string
	age  int
}

// time GOGC=50 ./pointers_ex_3
// GOGC=off makes program faster
// GOGC=100 is the default
//
// Run this to see the garbage collections
// time GODEBUG=gctrace=1 ./pointers_ex_3
func main() {
	slowProgram()
	fastProgram()
}

func slowProgram() {
	var people []Person
	// In this case there is a lot of memory being moved when the capiticy of the slice
	// is outgrown
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Jørgen", 31))
	}
}

func fastProgram() {
	// Allocate the memomry once for 10 000 000 persons
	people := make([]Person, 10_000_000)

	for i := 0; i < 10_000_000; i++ {
		people[i] = MakePerson("Jørgen", 31)
	}
}

func MakePerson(name string, age int) Person {
	return Person{
		age:  age,
		name: name,
	}
}
