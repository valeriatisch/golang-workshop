package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", uploadHandler)
	http.HandleFunc("/imgs/", listFilesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "templates/upload.gohtml", nil)
	} else if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("image")
		defer file.Close()
		if err != nil {
			http.Error(w, "Failed to retrieve image from form data", http.StatusBadRequest)
			return
		}

		// Create a new file on the server to save the uploaded image
		f, err := os.Create("imgs/" + handler.Filename)
		if err != nil {
			http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Copy the image file data to the newly created file
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, "Failed to save image file", http.StatusInternalServerError)
			return
		}

		renderTemplate(w, "templates/upload.gohtml", "Image uploaded successfully")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	dirPath := "./imgs"

	// Read the directory
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		println(err.Error())
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}

	// Create a slice of file names
	fileNames := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	// Render the template
	err = renderTemplate(w, "templates/list_files.gohtml", fileNames)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	return nil
}
