package main

/* ------------------------------------- *
 * Following the golang tutorial         *
 * https://golang.org/doc/articles/wiki/ *
 * ------------------------------------- */

import (
	"fmt"
	"html/template" //html templating
	"io/ioutil"
	"net/http"
)

// Make a page struct, has a title and body
type Page struct {
	Title string
	Body  []byte
}

// Saves a page in memory
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// return filename + data + permissions
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Save time writing code
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// Load a page to read contents
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}

	renderTemplate(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func main() {
	/*
		-- Demonstration of how page save/load works
		p1 := &Page{"TestPage", []byte("This is a sample")}
		p1.save()
		p2, _ := loadPage("TestPage")
		fmt.Println(string(p2.Body))
	*/

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	fmt.Println("Listening...")
	http.ListenAndServe(":8018", nil)
}
