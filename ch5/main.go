package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(addTo(0, 1, 2, 3))
	fmt.Println(addTo(10, 1, 2, 3))

	result, reminder, err := divAndReminder(5, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
	fmt.Println(reminder)
}

func addTo(base int, params ...int) []int {
	out := make([]int, 0, len(params))
	for _, v := range params {
		out = append(out, v+base)
	}
	return out
}

func divAndReminder(div, denom int) (result int, reminder int, err error) {
	if denom == 0 {
		return 0, 0, errors.New("cant divide by zero")
	}
	result = div / denom
	reminder = div % denom

	return result, reminder, nil
}
