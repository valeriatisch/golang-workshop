package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// A template is a string or file containing HTML & placeholders or loops enclosed in double braces {{ }}
// A template allows us to generate HTML dynamically.
// Template files typically end with .gohtml or .tmpl
var tmpl *template.Template

// FuncMap is a map of functions that can be used in a template.
var fm map[string]any = template.FuncMap{
	"upper":  strings.ToUpper,
	"double": double,
}

func init() {
	// Let's ensure that the template parsing and error handling are performed once in the beginning
	// Must() simplifies error handling as it panics when an error occurs.
	// tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	// Templates with functions
	// Funcs need to be registered before parsing
	tmpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

type Person struct {
	Name string
	Age  int
}

type Book struct {
	Title string
}

func main() {
	// Create a file called index.html
	file, err := os.Create("index.html")
	checkError(err)
	defer file.Close()

	// Placeholder {{ . }}
	err = tmpl.ExecuteTemplate(file, "placeholder.gohtml", "Jane")
	checkError(err)

	// Empty file
	file.Truncate(0)

	// Variables {{ $name := . }}
	err = tmpl.ExecuteTemplate(file, "variable.gohtml", "John")
	checkError(err)

	file.Truncate(0)

	// Structs {{ .Field }}
	p1 := Person{
		Name: "Emma",
		Age:  26,
	}
	err = tmpl.ExecuteTemplate(file, "structs.gohtml", p1)

	file.Truncate(0)

	// Loops {{ range . }}
	// Struct of slices
	p2 := Person{
		Name: "Tom",
		Age:  30,
	}
	b1 := Book{
		Title: "Harry Potter",
	}
	b2 := Book{
		Title: "Lord of the Rings",
	}
	data := struct {
		People []Person
		Books  []Book
	}{
		People: []Person{p1, p2},
		Books:  []Book{b1, b2},
	}
	err = tmpl.ExecuteTemplate(file, "loops.gohtml", data)
	file.Truncate(0)

	// Functions
	err = tmpl.ExecuteTemplate(file, "functions.gohtml", p1)

	// Look into the documentation for pre-defined global functions

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func double(x int) int {
	return x * 2
}
