package main

import "fmt"

func main() {

	m := Mananger{
		Employee: Employee{
			Name: "Bob Ross",
			ID:   "BOBBY",
		},
		Reports: []Employee{},
	}

	fmt.Printf(m.Description(), "from employee type")

	employees := m.FindNewEmployees()
	fmt.Println()
	for _, v := range employees {
		fmt.Printf(v.Name, "from manager type")
	}

	fmt.Println()
	fmt.Println()
	exampleWhenSameName()
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e, e.Name, e.ID)
}

type Mananger struct {
	Employee
	Reports []Employee
}

func (m Mananger) FindNewEmployees() []Employee {
	return []Employee{
		{"Jørgen", "JO"},
	}
}

type Inner struct {
	x int
}

type Outer struct {
	Inner
	x int
}

func exampleWhenSameName() {
	o := Outer{
		Inner: Inner{
			x: 10,
		},
		x: 20,
	}

	fmt.Println(o.x)
	// If the names of the field would have been unique
	// it could have been referenced via o.x
	fmt.Println(o.Inner.x)
}
