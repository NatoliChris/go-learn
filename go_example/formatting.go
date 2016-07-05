package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var p = fmt.Printf

func main() {
	a := point{1, 2}

	// verbs
	p("%v\n", a)

	// Shows struct names
	p("%+v\n", a)

	//Shows the Go representation
	p("%#v\n", a)

	//Shows the type
	p("%T\n", a)

	//Formatting Booleans
	p("%t\n", true)

	//Number representations
	p("%d\n", 123)
	p("%b\n", 123)         //Binary
	p("%c\n", 123)         //Character
	p("%x\n", 123)         //Hex
	p("%f\n", 12.3)        //floating point
	p("%e\n", 123400000.0) //scientific v1
	p("%E\n", 123400000.0) //scientific v2

	// String
	p("%s\n", "\"string\"")
	p("%q\n", "\"string\"") // Double quote

	// Misc
	p("%p\n", &a)
	p("|%6d|%6d|\n", 12, 1234)

	b := fmt.Sprintf("a %s", "string")
	fmt.Println(b)
	fmt.Fprintf(os.Stderr, "an %s\n", "Err")
}
