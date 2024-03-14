package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	// both these are equivelant
	zeroValueStruct()
	zeroValueStructLiteral()
	structLiteralNoNameParams()
	anonymousStructs()
}

func zeroValueStruct() {
	var jorgen person
	println("name", jorgen.name)
	println("age", jorgen.age)
	println("pet", jorgen.pet)
}

func zeroValueStructLiteral() {
	jorgen := person{}
	println("name", jorgen.name)
	println("age", jorgen.age)
	println("pet", jorgen.pet)
}

func structLiteralNoNameParams() {
	p := person{"jorgen", 30, "zero"}
	fmt.Println(p)

	fmt.Println("name", p.name)
	fmt.Println("age", p.age)
	fmt.Println("pet", p.pet)
}

func anonymousStructs() {
	var person struct {
		name string
		age  int
	}

	person.age = 30
	person.name = "Jonas"

	fmt.Println("age", person.age)
	fmt.Println("name", person.name)

	pet := struct {
		name string
	}{
		name: "Fido",
	}

	fmt.Println(pet)
}
