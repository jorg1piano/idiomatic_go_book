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
