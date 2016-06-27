package main

import "fmt"

func forloop() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func ifelse() {
	num := 9
	if num < 0 {
		fmt.Println("neg")
	} else if num > 10 {
		fmt.Println("gt 10")
	} else {
		fmt.Println("Your number is", num)
	}
}

func main() {
	fmt.Println("Running For Loops")
	forloop()

	fmt.Println("Running IF/ELSE")
	ifelse()
}
