package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Fred"})

	s := person{name: "Chris", age: 21}

	fmt.Println(s.name)
}
