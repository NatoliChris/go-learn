package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		upper := strings.ToUpper(scan.Text())

		fmt.Println(upper)
	}

	if err := scan.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}
