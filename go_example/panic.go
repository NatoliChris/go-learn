package main

import "os"

func main() {
	// Typically means something went wrong
	panic("A problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

}
