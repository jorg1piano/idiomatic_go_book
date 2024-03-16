package main

import "fmt"

type LogicProvider struct{}

// Value reciever for LogicProvider
func (lp LogicProvider) Process(data string) string {
	for i := 0; i < 10; i++ {
		fmt.Println(data)
	}
	return fmt.Sprintf("Printed %s 10 times", data)
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() string {
	data := "Hello, World!"
	return c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}

	res := c.Program()

	fmt.Println(res)
}
