package main

import (
	"fmt"
	"io"
	"net/http"
)

// will echo back
func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/hello", hello)
	fmt.Println("Listening...")
	http.ListenAndServe(":8018", nil)
}
