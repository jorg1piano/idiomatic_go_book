package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	people := []person{
		{"Janusrgen", "Andersen", 32},
		{"Jonas", "Andersen", 32},
		{"Tina", "Mathisen", 30},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	fmt.Println(people)
}
