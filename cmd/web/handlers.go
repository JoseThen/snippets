package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a handler function for the homepage that writes "Hello from Snippetbox"
func home(w http.ResponseWriter, r *http.Request) {
	// Because the "/" path is a catchall and will route here, we need to check if the request is for the homepage
	if r.URL.Path != "/" {
		// If it isn't the homepage, then return a 404 not found error
		http.NotFound(w, r)
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
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Execute the template and if there is an error return with a 500 error
	// In other words return/run the template and if there is an error return a 500 error
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Fetch the ID from the URL query param and if its less than 1 or not a number, return a 400 Bad Request error
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Write the snippet text to the response
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != "POST" {
		// If not a post, respond with which methods are allowed
		w.Header().Set("Allow", http.MethodPost)
		// Return a 405 Method Not Allowed error
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
