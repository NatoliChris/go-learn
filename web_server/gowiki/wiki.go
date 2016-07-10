package main

/* ------------------------------------- *
 * Following the golang tutorial         *
 * https://golang.org/doc/articles/wiki/ *
 * ------------------------------------- */

import (
	"fmt"
	"io/ioutil"
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

// Load a page to read contents
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{"TestPage", []byte("This is a sample")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
