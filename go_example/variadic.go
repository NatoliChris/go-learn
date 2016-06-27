package main

import "fmt"

func sum(vals ...int) {
	fmt.Print(vals, " ")
	total := 0
	for _, num := range vals {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
