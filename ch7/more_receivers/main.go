package main

import (
	"fmt"
	"time"
)

func main() {
	var c Counter
	c.Increment()
	c.Increment()
	c.Increment()
	fmt.Println(c.String())

	doUpdateWrong(c)
	fmt.Println(c.String(), "outside")
	doUpdateRight(&c)
	fmt.Println(c.String(), "outside")
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++ // zero value is 0 so this is fine
	c.lastUpdated = time.Now()
}

func (c *Counter) String() string {
	return fmt.Sprintf("total: %d time: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	// When exiting this function the counter on the struct passed in will not be changed.
	// Struct are also copied when passed as arguments so any changes to the copy wont affect
	// the value passed in.
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}
