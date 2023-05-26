package main

import (
	"log"
	"text/template"
)

// A template is a string or file containing HTML & placeholders or loops enclosed in double braces {{ }}
// A template allows us to generate HTML dynamically.
// Template files typically end with .gohtml or .tmpl
var tmpl *template.Template

// FuncMap is a map of functions that can be used in a template.

func init() {
	// Let's ensure that the template parsing and error handling are performed once in the beginning
	// Must() simplifies error handling as it panics when an error occurs.

	// Templates with functions
	// Funcs need to be registered before parsing

}

type Person struct {
	Name string
	Age  int
}

type Book struct {
	Title string
}

func main() {
	// TODO: Create a file called index.html

	// Placeholder {{ . }}

	// Empty file

	// Variables {{ $name := . }}

	// Structs {{ .Field }}

	// Loops {{ range . }}
	// Struct of slices

	// Functions

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
