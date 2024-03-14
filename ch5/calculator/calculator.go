package main

import (
	"fmt"
	"strconv"
)

type operationsTable struct {
	a, operation, b string
}

type mathOperation func(a, b int) int

var opMap = map[string]mathOperation{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func main() {
	operations := []operationsTable{
		{"1", "+", "2"},
		{"1", "*", "2"},
		{"1", "/", "2"},
		{"1", "-", "2"},
	}

	calculate(operations)
}

func calculate(operations []operationsTable) {
	for _, o := range operations {
		f, ok := opMap[o.operation]
		if !ok {
			fmt.Println("Invalid operation", o.operation)
			continue
		}
		a, err := strconv.Atoi(o.a)
		if err != nil {
			fmt.Println(err)
			continue
		}

		b, err := strconv.Atoi(o.b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := f(a, b)
		fmt.Println("Result", result)
	}
}
