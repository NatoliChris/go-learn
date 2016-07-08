package main

import (
	"fmt"
	"strconv"
)

func main() {

	// parseFloat(number, precision)
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// the 0 means base, 64 is bit size
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// Recognizes hex
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//simple base 10
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("errror")
	fmt.Println(e)
}
