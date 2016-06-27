package main

import "fmt"

func main() {
	/*
		Initializing array
		Syntax:
		var _____ [size]type
	*/

	var a [5]int
	fmt.Println("Initialized: ", a)

	//Assigning like normal arrays, or can assign on initialization
	a[4] = 100
	fmt.Println("Assigned: ", a)
	fmt.Println("Length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Declare + Initialize: ", b)

	/* Slices */
	d := make([]string, 3)
	d[0] = "a"
	d[1] = "b"
	d[2] = "c"
	sl := d[0:2]

	fmt.Println("Slice: ", sl)

	// Trying TwoDimension

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner := i + 1
		twoD[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("Two Dimension: ", twoD)
}
