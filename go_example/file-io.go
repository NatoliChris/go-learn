package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	/* ---------------------
			File Reading
	   --------------------- */
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Go to known location in file
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// more robust read
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// No built in rewind but seek(0,0)
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 Bytes: %s\n", string(b4))

	// Close the file!
	f.Close()
	/* ---------------------
			File Writing
	   --------------------- */

	d1 := []byte("Hello\ngo\n")
	// Dump bytes into a file with permissions
	ex := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(ex)

	ofile, err := os.Create("/tmp/dat2")
	check(err)

	defer ofile.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n5, err := ofile.Write(d2)
	check(err)
	fmt.Printf("wrote: %d bytes\n", n5)

	n6, err := ofile.WriteString("Writes\n")
	fmt.Printf("wrote %d bytes\n", n6)

	// Flushes the file to writable storage
	ofile.Sync()

	w := bufio.NewWriter(ofile)
	n7, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n7)

	w.Flush()
}
