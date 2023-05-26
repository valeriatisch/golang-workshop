package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Structure to hold book data
type Book struct {
	ID     int
	Title  string
	Author string
}

// Initialize some book data
var books = []Book{
	{ID: 1, Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: 2, Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
}

var tmpl *template.Template

func init() {
	// Parse all templates
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// Register the functions to handle requests to "/books" & "/createbook"

	// Start the HTTP server
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleBooks(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Requested path:", req.URL.Path)
	// Determine the HTTP method
	switch req.Method {
	// If it's a GET request
	case http.MethodGet:
		// Trim "/books" prefix from the URL path
		path := strings.TrimPrefix(req.URL.Path, "/books")
		switch {
		// If the path is "/books" or "/books/", list all books
		case path == "" || path == "/": // for example, "/books/1" => "/1"
			// If the path is "/books" or "/books/", list all books
			err := tmpl.ExecuteTemplate(resp, "books.gohtml", books)
			if err != nil {
				http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		// If the path is "/books/[some number]", show the specific book
		case strings.HasPrefix(path, "/"):
			// Remove leading slash if there is one
			path = strings.TrimSuffix(path, "/")
			// If the ID is not a valid number, return a 404 Not Found response
			id, err := strconv.Atoi(strings.TrimPrefix(path, "/"))
			if err != nil {
				http.NotFound(resp, req)
				return
			}
			// Loop through the books and find the one with the given ID
			// If the book exists, show its details = execute the "book.gohtml" template
			for _, book := range books {
				if book.ID == id {
					err = tmpl.ExecuteTemplate(resp, "book.gohtml", book)
					if err != nil {
						http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
					}
					return
				}
			}
			// If no book with the given ID exists, return a 404 Not Found response
			http.NotFound(resp, req)
		default:
			// If the path doesn't match any of the above cases, return a 404 Not Found response
			http.NotFound(resp, req)
		}
	// If it's a POST request
	case http.MethodPost:
		// Return a 404 is the path is not "/books" or "/books/"
		if req.URL.Path != "/books" && req.URL.Path != "/books/" {
			http.NotFound(resp, req)
			return
		}

		// Parse the form data with req.ParseForm()
		if err := req.ParseForm(); err != nil {
			http.Error(resp, "Failed to parse form", http.StatusBadRequest)
			return
		}
		fmt.Println("Requested data parsed to form:", req.Form)

		// Put data from req.FormValue() into a Book struct
		book := Book{
			ID:     len(books) + 1,
			Title:  req.FormValue("title"),
			Author: req.FormValue("author"),
		}
		// Append new book to books slice
		books = append(books, book)

		// Redirect to the book list
		http.Redirect(resp, req, "/books", http.StatusSeeOther)
	default:
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createBook(resp http.ResponseWriter, req *http.Request) {
	// Parse the form data with req.ParseForm()

	// Put data from req.FormValue() into a Book struct

	// Append new book to the books slice

	// Redirect to the book list

}
