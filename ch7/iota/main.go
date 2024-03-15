package main

import "fmt"

const (
	Harmony = iota
	Chord
	Interval
	Scale
)

func main() {
	fmt.Println("Chord", Chord)
	fmt.Println("Interval", Interval)
	fmt.Println("Scale", Scale)
}
