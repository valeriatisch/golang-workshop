package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	templates = template.Must(template.ParseFiles("templates/submit.gohtml", "templates/welcome.gohtml"))
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if mycookie exists
	// If the cookie doesn't exist, display submit.gohtml with renderTemplate()
	// If the cookie exists, redirect to the welcome page

}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form values

	// Generate a UUID for the user

	// Create a new cookie with the UUID as the value

	// Set the cookie

	// Redirect to the welcome page

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the cookie exists
	cookie, err := r.Cookie("mycookie")
	if err == http.ErrNoCookie {
		// If the cookie doesn't exist, redirect to the home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if err != nil {
		// Handle other errors
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the uuid from the cookie
	uid := cookie.Value

	// Get the visit count from the cookie
	count := 1
	countCookie, err := r.Cookie("visitcount")
	if err == http.ErrNoCookie {
		// If the cookie doesn't exist, create a new one
		countCookie = &http.Cookie{
			Name:  "visitcount",
			Value: strconv.Itoa(count),
			Path:  "/",
		}
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		// Update cookie value
		count, _ = strconv.Atoi(countCookie.Value)
		countCookie.Value = strconv.Itoa(count + 1)
	}

	http.SetCookie(w, countCookie)

	// Data for rendering the welcome template
	data := struct {
		Uid   string
		Count int
	}{
		Uid:   uid,
		Count: count,
	}

	// Render the welcome template
	renderTemplate(w, "welcome.gohtml", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
