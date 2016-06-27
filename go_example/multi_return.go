package main

import "fmt"

func vals() (int, int) {
	return 5, 3
}

func main() {
	fmt.Println(vals())

	a, b := vals()
	fmt.Printf("a: %d, b: %d\n", a, b)
}
