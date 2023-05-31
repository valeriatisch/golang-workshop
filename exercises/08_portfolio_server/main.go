package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Create a Network struct with the fields: Name and Link, add json tags
type Network struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// TODO: Create a Portfolio struct with the fields:
// Name, Title, About, Networks, PhotoPath, PhotoFile (of type multipart.File)
// TODO: Add json tags
// PhotoPath should have the tag: `json:"-"` to ignore them later when marshalling to JSON
type Portfolio struct {
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	About     string    `json:"about"`
	Networks  []Network `json:"networks"`
	PhotoPath string    `json:"-"`
	PhotoFile multipart.File
}

// TODO: Create a user struct with the fields: Username and Portfolio
type User struct {
	Username  string    `json:"username"`
	Portfolio Portfolio `json:"portfolio"`
}

// TODO: Create a *template.Template variable
var tmpl *template.Template

func init() {
	// TODO: Parse all templates in the templates folder
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	// TODO: Register the handler functions
	http.HandleFunc("/", createUserHandler)
	http.HandleFunc("/download", downloadHandler)

	// TODO: Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// TODO: Execute the create.gohtml template
		tmpl.ExecuteTemplate(w, "create.gohtml", nil)
	} else if r.Method == http.MethodPost {
		// TODO: Parse the form values into a Portfolio struct
		// The formValue keys are named as follows: name, title, about, networkName, networkLink, username
		portfolio := Portfolio{
			Name:  r.FormValue("name"),
			Title: r.FormValue("title"),
			About: r.FormValue("about"),
		}

		portfolio.Networks = make([]Network, 1)
		portfolio.Networks[0].Name = r.FormValue("networkName")
		portfolio.Networks[0].Link = r.FormValue("networkLink")

		// TODO: Process the uploaded photo file with .FormFile()
		// TODO: Save the photo to /uploads and its data to the Portfolio struct (PhotoFile, PhotoPath)
		photo, handler, err := r.FormFile("photo")
		if err == nil {
			defer photo.Close()
			portfolio.PhotoFile = photo

			filename := handler.Filename
			portfolio.PhotoPath = filename
			err = saveUploadedFile(photo, filepath.Join("uploads", filename))
			if err != nil {
				http.Error(w, "Failed to save photo file", http.StatusInternalServerError)
				return
			}
		}

		// TODO: Parse the username and portfolio into a User struct
		user := User{
			Username:  r.FormValue("username"),
			Portfolio: portfolio,
		}

		fmt.Println(user)

		// TODO: Generate the HTML file for the user's portfolio with generateHTMLFile()
		err = generateHTMLFile(user)
		if err != nil {
			http.Error(w, "Failed to generate HTML file", http.StatusInternalServerError)
			return
		}

		// TODO: Redirect to /download
		http.Redirect(w, r, "/download", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=portfolio.html")
	// TODO: Serve the portfolio.html file with http.ServeFile()
	http.ServeFile(w, r, "portfolio.html")
}

func generateHTMLFile(user User) error {
	// TODO: Create the portfolio.html file
	file, err := os.Create("portfolio.html")
	checkError(err)
	defer file.Close()

	// Update the photo path in the portfolio HTML template
	portfolio := user.Portfolio
	if portfolio.PhotoPath != "" {
		portfolio.PhotoPath = strings.Replace(portfolio.PhotoPath, `\`, "/", -1)
	}

	// TODO: Execute the template with the user's portfolio
	err = tmpl.ExecuteTemplate(file, "template.gohtml", portfolio)
	checkError(err)

	return nil
}

func saveUploadedFile(file multipart.File, path string) error {
	out, err := os.Create(path)
	checkError(err)
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
