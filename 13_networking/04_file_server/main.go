package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)

	// Serve about.html at the "/about" path
	http.HandleFunc("/about", serveAbout)

	// Serve static files (including styles.css) under the "/static/" path prefix
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	// Serve about.html file
	http.ServeFile(w, r, "static/about.html")
}
