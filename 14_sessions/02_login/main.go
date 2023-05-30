package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	templates *template.Template
	users     = make(map[string]string)
	sessions  = make(map[string]Session)
)

type User struct {
	Username string
	Password []byte
}

type Session struct {
	Username  string
	SessionID string
}

func init() {
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
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
	if isAuthenticated(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

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
			Username:  username,
			SessionID: sessionID,
		}

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
	session := sessions[sessionID]

	renderTemplate(w, "home.gohtml", session.Username)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := getSessionID(r)

	delete(sessions, sessionID)

	clearSessionCookie(w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Helper Functions
func setSessionCookie(w http.ResponseWriter, sessionID string) {
	cookie := http.Cookie{
		Name:  "session",
		Value: sessionID,
		Path:  "/",
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

func isAuthenticated(r *http.Request) bool {
	sessionID := getSessionID(r)

	_, ok := sessions[sessionID]
	return ok
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
