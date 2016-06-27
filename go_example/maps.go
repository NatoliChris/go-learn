package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 3

	fmt.Println("Map: ", m)

	v1 := m["k1"]
	fmt.Println("value m[k1]:", v1)
	fmt.Println("Length: ", len(m))

	delete(m, "k2")
	fmt.Println("After `delete(m, \"k2\"):`", m)

}
