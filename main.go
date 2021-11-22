package main

import (
	"log"
	"net/http"
)

// Define a handler function for the homepage that writes "Hello from Snippetbox"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Use the home function as the handler for the "/" path
	mux.HandleFunc("/", home)

	log.Println("Listening on :4000")
	// Use http.ListenAndServe to start the server on the specified port using our created mux.
	// If it returns an error then log it and exit the program.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
