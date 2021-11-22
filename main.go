package main

import (
	"fmt"
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
	w.Write([]byte("Hello from Snippetbox"))
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

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Use the home function as the handler for the "/" path
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Listening on :4000")
	// Use http.ListenAndServe to start the server on the specified port using our created mux.
	// If it returns an error then log it and exit the program.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
