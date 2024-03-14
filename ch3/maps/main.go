package main

import "fmt"

func main() {
	nilMapExample()
	mapLiteralExample()
	mapLenFunc()
	readingAndWrintingAMap()
	commaOkayIdiom()
}
func nilMapExample() {
	var nilMap map[string]int
	fmt.Println(nilMap)
	// returns the zero value
	fmt.Println(nilMap["hello"])

	// nilMap["hello"] = 69
	// output
	// panic: assignment to entry in nil map
}

func mapLiteralExample() {
	namesInPlaces := map[string][]string{
		"Selkop": {"Asbjørn", "Torild"},
		"Alta":   {"Jorgen", "Jonas"},
	}

	fmt.Println(namesInPlaces)
}

func mapLenFunc() {
	m := make(map[int]int, 3)

	m[2] = 4
	m[4] = 6
	m[8] = 10
	m[10] = 12

	fmt.Println(len(m)) // output 4
}

func readingAndWrintingAMap() {
	fmt.Println()
	fmt.Println("readingAndWrintingAMap")
	totalWins := map[string]int{}
	totalWins["Orca"] = 1
	totalWins["Lions"] = 1

	fmt.Println(totalWins["Orca"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 10
	fmt.Println(totalWins["Lions"])
}

func commaOkayIdiom() {
	m := map[string]int{
		"Jørgen": 31,
		"Jonas":  31,
	}

	v, ok := m["Jørgen"]
	fmt.Println(v, ok)

	v, ok = m["Jonas"]
	fmt.Println(v, ok)

	v, ok = m["Marlene"]
	fmt.Println(v, ok)
}
