package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	st := sort.StringsAreSorted(strs)
	fmt.Println("String sorted: ", st)
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	st = sort.StringsAreSorted(strs)
	fmt.Println("String sorted: ", st)
	//sorts specify the type
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	//Checks if the given array is sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
