package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum: ", sum)

	fmt.Println("Using for key, value in range of map")

	kvs := map[string]string{"A": "apple", "B": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}

	for i, c := range "Hello" {
		fmt.Println(i, c)
	}
}
