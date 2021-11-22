package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("Show a specific snippet"))
}

// createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
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
