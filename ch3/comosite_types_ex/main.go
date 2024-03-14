package main

import "fmt"

func main() {
	var greetings []string = []string{
		"Hello", "Hi", "Hola", "Halloen",
	}

	a := greetings[:2]
	b := greetings[1:]
	fmt.Println(a)
	fmt.Println(b)

	examplesOfRune()
}

func examplesOfRune() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))

	fmt.Println("TEST")
	// fmt.Println(len(c))
	fmt.Printf("%T", rune('T'))
	fmt.Println()
	fmt.Printf("%T", rune('👩'))
	fmt.Println()
	fmt.Printf("%c", byte(65))
	fmt.Println()
	fmt.Printf("%T", byte(65))
	fmt.Println()
}
