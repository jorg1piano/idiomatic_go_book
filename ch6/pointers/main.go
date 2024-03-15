package main

import "fmt"

func main() {
	fmt.Println("Hello pointers")
	dereferenceAndIndirectionOperator()
	fmt.Println()

	nilPointers()
	pointWithNew()
	pointerRepitition()
	structPointers()
	playingWithPointerParams()
}

func dereferenceAndIndirectionOperator() {
	x := 10
	pointerX := &x
	valueX := *pointerX
	fmt.Println(pointerX)
	fmt.Println(valueX)
}

func nilPointers() {

	var nilPointer *int

	fmt.Println(nilPointer)

	// Dereferencing a nil pointer will cause a runtime error
	// fmt.Println(*nilPointer)
	// output:
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x1045a8b38]
}

func pointWithNew() {
	var x *int = new(int)
	fmt.Println(*x, "*x should be 0")
	var y = new(int)
	fmt.Println(*y, "*y should be 0")

	var str = new(string)
	fmt.Println(*str, "is an empty string that points to", str)
}

func pointerRepitition() {
	x := 69

	pointerX := &x

	fmt.Println(x, *pointerX, "is the same")
	fmt.Println(&x, pointerX, "is the same")
}

func structPointers() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	// p := person{
	// 	FirstName:  "John",
	// 	MiddleName: "Perry",
	// 	LastName:   "Doe",
	// }
	// output:
	// cannot use "Perry" (untyped string constant) as *string value in struct literalcompilerIncompatibleAssign

	var perry = "Perry"
	p := person{
		FirstName:  "John",
		MiddleName: &perry,
		LastName:   "Doe",
	}

	fmt.Println(p.FirstName, *p.MiddleName, p.LastName)

	// Workaround with generic function
	p2 := person{
		FirstName:  "John",
		MiddleName: makePointer("Perry"),
		LastName:   "Doe",
	}

	fmt.Println(p2.FirstName, *p2.MiddleName, p2.LastName)
}

// This is a helper function that takes a constant value and turns it into a pointer
func makePointer[T any](t T) *T {
	return &t
}

// Some notes on pointers in g0
// If you pass a pointer to a methode, the parameter is still copied.
// We are copying the pointer, but the copied pointer still points to the same data
//

func playingWithPointerParams() {
	// There is a godcha, if you pass a nil pointer to a function, you cannot
	// make the value non-nil. You can't reassign the value unless the pointer was not nil.
	fmt.Println("failedUpdate")
	var f *int
	failedUpdate(f)
	fmt.Println(f)

	x := 10
	actualUpdate(&x)
	fmt.Println(x)

	px := 0
	fmt.Println(&px, "px address in decleared scope")
	updateTo20(&px)
	fmt.Println(px)
}

func failedUpdate(g *int) {
	x := 10
	// The copy of the pointer now points to the new address
	g = &x

	// Updating this pointer will not change the original pointer
	*g = 420
}

func actualUpdate(g *int) {
	// Change the value g points to
	*g = 69
}

func updateTo20(px *int) {
	fmt.Println(&px, "px address in function scope")
	*px = 20
}
