package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello, World!")
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Goodbye")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", goodbye)
	fmt.Println("Listening...")
	http.ListenAndServe(":8018", nil)
}
