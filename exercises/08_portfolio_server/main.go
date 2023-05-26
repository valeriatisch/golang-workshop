package main

import (
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

// TODO: Create a Network struct with the fields: Name and Link, add json tags


// TODO: Create a Portfolio struct with the fields:
// Name, Title, About, Networks, PhotoPath, PhotoFile (of type multipart.File from the mime/multipart package)
// TODO: Add json tags
// PhotoPath should have the tag: `json:"-"` to ignore them later when marshalling to JSON

// TODO: Create a user struct with the fields: Username and Portfolio

// TODO: Create a *template.Template variable

func init() {
	// TODO: Parse all templates in the templates folder

}

func main() {
	// TODO: Register the handler functions

	// TODO: Start the server on port 8080
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// TODO: Execute the create.gohtml template

		} else if r.Method == http.MethodPost {
		// TODO: Parse the form values into a Portfolio struct
		// The formValue keys are named as follows: name, title, about, networkName, networkLink, username

		// TODO: Process the uploaded photo file with .FormFile()
		// TODO: Save the photo to /uploads and its data to the Portfolio struct (PhotoFile, PhotoPath)

		// TODO: Parse the username and portfolio into a User struct

		// TODO: Generate the HTML file for the user's portfolio with generateHTMLFile()

		// TODO: Redirect to /download

	} else {
		// TODO: Write a MethodNotAllowed error to the response

	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=portfolio.html")
	// TODO: Serve the portfolio.html file with http.ServeFile()

}

func generateHTMLFile(user User) error {
	// TODO: Create the portfolio.html file

	// Update the photo path in the portfolio HTML template
	portfolio := user.Portfolio
	if portfolio.PhotoPath != "" {
		portfolio.PhotoPath = strings.Replace(portfolio.PhotoPath, `\`, "/", -1)
	}

	// TODO: Execute the template with the user's portfolio and write it to the file

	return nil
}
