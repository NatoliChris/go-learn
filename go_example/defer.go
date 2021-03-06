package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	// says to leave function until after
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	//write line to file
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	f.Close()
}
