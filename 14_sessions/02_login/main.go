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
		username := r.FormValue("username") // retrieve username
		password := r.FormValue("password") // retrieve password

		// check if username already exists
		if _, ok := users[username]; ok {
			fmt.Fprintln(w, "Username already taken. Please choose a different username.")
			return
		}

		// encrypt password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// save new username and hashed password
		users[username] = string(hashedPassword)

		// redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	renderTemplate(w, "signup.gohtml", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {// login handler gets a response writer which is used to construct an HTTP response, and r and http request
	if r.Method == "POST" { // check whether request method is POST
		username := r.FormValue("username") // save username form value in variable username
		password := r.FormValue("password") // save passowrd form value in variable password

		hashedPassword, ok := users[username] // from users map get hashedpassword value
		if !ok { // check if retrieving hashedPassword is not ok
			fmt.Fprintln(w, "Invalid username. Please try again.") // print inavalid username..
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) // compare hash with unhashed password from form
		if err != nil { //check for error
			fmt.Fprintln(w, "Invalid password. Please try again.")
			return
		}

		sessionID := uuid.New().String() // create new session id

		sessions[sessionID] = Session{ //
			Username:  username,
			SessionID: sessionID,
		}

		setSessionCookie(w, sessionID) // add SessionID to http response

		http.Redirect(w, r, "/home", http.StatusSeeOther) // redirects user to url + /home
		return
	}

	renderTemplate(w, "login.gohtml", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//check if the user if authenticated, if not redirect to the login page
	if !isAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
// retrive the session id and put the id to the map
	sessionID := getSessionID(r)
	session := sessions[sessionID]

	renderTemplate(w, "home.gohtml", session.Username)
}

// Handles requests to the /logout URL => Ends the current session
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieves the users sessionID from a cookie if it exists
	sessionID := getSessionID(r)

	// Deletes the session of the user from the map
	delete(sessions, sessionID)

	// Clears the users session cookie (sets it to an expired cookie)
	clearSessionCookie(w)

	// Redirects the user to the "/" (root) page
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
