package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a handler function for the homepage that writes "Hello from Snippetbox"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because the "/" path is a catchall and will route here, we need to check if the request is for the homepage
	if r.URL.Path != "/" {
		// If it isn't the homepage, then return a 404 not found error
		app.notFound(w)
		return
	}

	// Setup a slice with the paths to your files needed. the homepage must be first
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles function to parse the files slice
	// if there is an error return with a 500 error
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	// Execute the template and if there is an error return with a 500 error
	// In other words return/run the template and if there is an error return a 500 error
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// showSnippet handler function
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Fetch the ID from the URL query param and if its less than 1 or not a number, return a 400 Bad Request error
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	// Write the snippet text to the response
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// createSnippet handler function
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != "POST" {
		// If not a post, respond with which methods are allowed
		w.Header().Set("Allow", http.MethodPost)
		// Return a 405 Method Not Allowed error
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
