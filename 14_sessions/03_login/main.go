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

// TODO:
// 1. Find a partner & try to understand the code
// 2. Partner A should read & comment the functions createUserDB(), indexHandler(), signupHandler()
// 3. Partner B should read & comment loginHandler(), homeHandler(), logoutHandler() 
// 4. Explain the code to each other

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
	// Check if the user is authenticated
	if isAuthenticated(r) {
		// Redirect to the home page
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	// Render the index page with sign up and login buttons
	renderTemplate(w, "index.gohtml", nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if _, ok := users[username]; ok {
			fmt.Fprintln(w, "Username already taken. Please choose a different username.")
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		users[username] = string(hashedPassword)

		http.Redirect(w, r, "/login", http.StatusSeeOther)

		file, err := os.OpenFile("user.csv", os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		record := []string{username, string(hashedPassword)}
		writer.Write(record)
		writer.Flush()

		if err := writer.Error(); err != nil {
			log.Fatal(err)
		}
		return
	}

	renderTemplate(w, "signup.gohtml", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		hashedPassword, ok := users[username]
		if !ok {
			fmt.Fprintln(w, "Invalid username. Please try again.")
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			fmt.Fprintln(w, "Invalid password. Please try again.")
			return
		}

		sessionID := uuid.New().String()

		sessions[sessionID] = Session{
			Username:     username,
			SessionID:    sessionID,
			CreationTime: time.Now(),
		}

		counts[username]++

		setSessionCookie(w, sessionID)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "login.gohtml", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sessionID := getSessionID(r)

	session, ok := sessions[sessionID]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if sessionHasExpired(session) {
		delete(sessions, sessionID)
		clearSessionCookie(w)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := session.Username
	count := counts[username]
	counts[username]++

	context := struct {
		Username string
		Count    int
	}{
		Username: username,
		Count:    count,
	}

	renderTemplate(w, "home.gohtml", context)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)

	delete(sessions, sessionID)

	clearSessionCookie(w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Helper Functions

func setSessionCookie(w http.ResponseWriter, sessionID string) {
	cookie := http.Cookie{
		Name:    "session",
		Value:   sessionID,
		Expires: time.Now().Add(5 * time.Minute),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
}

func getSessionID(r *http.Request) string {
	cookie, err := r.Cookie("session")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func clearSessionCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
}

func sessionHasExpired(session Session) bool {
	duration := time.Since(session.CreationTime)

	return duration.Minutes() >= 5
}

func isAuthenticated(r *http.Request) bool {
	sessionID := getSessionID(r)

	_, ok := sessions[sessionID]
	return ok
}

func createUserDB() {
	file, err := os.OpenFile("user.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(records) == 0 {
		records = append(records, []string{"username", "password_hash"})

		writer := csv.NewWriter(file)
		writer.WriteAll(records)
		writer.Flush()

		if err := writer.Error(); err != nil {
			log.Fatal(err)
		}
	}

	for _, record := range records[1:] {
		username := record[0]
		passwordHash := record[1]

		users[username] = passwordHash
	}

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
