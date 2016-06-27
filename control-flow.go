package main

import (
	"fmt"
	"time"
)

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

func learnSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend :), today is:", time.Now().Weekday())
	default:
		fmt.Println("It's a weekday :(, it's only", time.Now().Weekday())
	}
}

func main() {

	//Learn how to do for loops
	fmt.Println("Running For Loops")
	forloop()

	//Simple If/Else ... basic language syntax
	fmt.Println("Running IF/ELSE")
	ifelse()

	fmt.Println("Switch Statements")
	learnSwitch()

}
