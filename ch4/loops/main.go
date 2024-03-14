package main

import "fmt"

func main() {
	fmt.Println("i := 0; i < 3; i++")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("; i< 3 ; i++")
	i := 0
	for ; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("i := 0; i < 3")
	for i := 0; i < 3; {
		fmt.Println(i)
		i++
	}

	// x := 0
	// for {
	// 	x++
	// 	fmt.Println("forever loop", x)
	// }

	fizzBuzz()
	prinOutNames()
	printOutFunNumbers()
	printScale()
	itterateOverStrings()
	fizzBuzzSwitch()
}

func fizzBuzz() {
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz", i)
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz", i)
			continue
		}
	}
}

func fizzBuzzSwitch() {
	for i := 0; i <= 15; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz", i)
		case i%3 == 0:
			fmt.Println("Fizz", i)
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func prinOutNames() {
	names := []string{"Jørgen", "Jonas", "Marlene"}

	for _, name := range names {
		fmt.Println(name)
	}

}

func printOutFunNumbers() {
	numbers := []int{69, 420, 1337}

	for i, v := range numbers {
		fmt.Println(i, v)
	}
}

func printScale() {
	scale := []string{"A", "B", "C", "F", "G"}

	for i, v := range scale {
		fmt.Println(i, v)
	}
}

func itterateOverStrings() {
	words := []string{"hello", "world"}

	for _, word := range words {
		for i, r := range word {
			// r is a rune
			fmt.Println(i, r, string(r))
		}
	}
}
