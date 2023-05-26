package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	templates *template.Template
	users     = make(map[string]string)
	sessions  = make(map[string]Session)
	counts    = make(map[string]int)
)

type User struct {
	Username string
	Password []byte
}

type Session struct {
	Username     string
	SessionID    string
	CreationTime time.Time
}

func init() {
	// Load the HTML templates
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	createUserDB()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/logout", logoutHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler Functions

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Check if the user is authenticated
	// If so, redirect to the /home page
	// If the user is not authenticated, render the index.gohtml page with sign up and login buttons
	
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Retrieve the form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if the username is already taken
		if _, ok := users[username]; ok {
			fmt.Fprintln(w, "Username already taken. Please choose a different username.")
			return
		}

		// Hash the password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Store the username and hashed password in the users map
		users[username] = string(hashedPassword)

		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)

		// Open the user.csv file in write mode and append
		file, err := os.OpenFile("user.csv", os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Create a new CSV writer
		writer := csv.NewWriter(file)

		// Add the username and password hash
		record := []string{username, string(hashedPassword)}

		// Write the record to the CSV file
		writer.Write(record)
		writer.Flush()

		if err := writer.Error(); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Render the signup page
	renderTemplate(w, "signup.gohtml", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Retrieve the form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if the username exists in the users map
		hashedPassword, ok := users[username]
		if !ok {
			fmt.Fprintln(w, "Invalid username. Please try again.")
			return
		}

		// Compare the hashed password with the provided password
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			fmt.Fprintln(w, "Invalid password. Please try again.")
			return
		}

		// Create a new session ID
		sessionID := uuid.New().String()

		// Store the session ID and session details in the sessions map
		sessions[sessionID] = Session{
			Username:     username,
			SessionID:    sessionID,
			CreationTime: time.Now(),
		}

		// Increment the visit count for the user
		counts[username]++

		// Set the session ID as a cookie
		setSessionCookie(w, sessionID)

		// Redirect to the home page
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	// Render the login page
	renderTemplate(w, "login.gohtml", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is authenticated
	if !isAuthenticated(r) {
		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get the session ID from the cookie
	sessionID := getSessionID(r)

	// Get the session details associated with the session ID
	session, ok := sessions[sessionID]
	if !ok {
		// If the session ID is not found, redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if the session has expired
	if sessionHasExpired(session) {
		// If the session has expired, delete the session and clear the session cookie
		delete(sessions, sessionID)
		clearSessionCookie(w)

		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get the username and visit count for the user
	username := session.Username
	count := counts[username]
	counts[username]++

	// Create a context for rendering the template
	context := struct {
		Username string
		Count    int
	}{
		Username: username,
		Count:    count,
	}

	// Render the home page with the context
	renderTemplate(w, "home.gohtml", context)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	sessionID := getSessionID(r)

	// Delete the session ID from the sessions map
	delete(sessions, sessionID)

	// Clear the session cookie by setting an expired cookie
	clearSessionCookie(w)

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Helper Functions

func setSessionCookie(w http.ResponseWriter, sessionID string) {
	// TODO: Create & set a session cookie with sessionID as Value & 5 minute expiration (hint: Expires field)
}

func getSessionID(r *http.Request) string {
	// TOOD: Get the sessionID from the session cookie
	return ""
}

func clearSessionCookie(w http.ResponseWriter) {
	// TODO: Clear the session by setting a cookie with an empty string for Value an an expiration time in the past

}

func sessionHasExpired(session Session) bool {
	// Calculate the duration since session creation
	duration := time.Since(session.CreationTime)

	// Check if the duration has exceeded the session expiration time
	return duration.Minutes() >= 5
}

func isAuthenticated(r *http.Request) bool {
	// Get the session ID from the cookie
	sessionID := getSessionID(r)

	// Check if the session ID exists in the sessions map
	_, ok := sessions[sessionID]
	return ok
}

func createUserDB() {
	// Open the user.csv file
	file, err := os.OpenFile("user.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the file is empty
	if len(records) == 0 {
		// If the file is empty, create a new header
		records = append(records, []string{"username", "password_hash"})

		// Write the header to the CSV file
		writer := csv.NewWriter(file)
		writer.WriteAll(records)
		writer.Flush()

		if err := writer.Error(); err != nil {
			log.Fatal(err)
		}
	}

	// Iterate over the records and populate the users map
	for _, record := range records[1:] {
		username := record[0]
		passwordHash := record[1]

		// Store the username and password hash in the users map
		users[username] = passwordHash
	}

	// Print the users map
	for username, passwordHash := range users {
		fmt.Printf("Username: %s, Password Hash: %s\n", username, passwordHash)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
