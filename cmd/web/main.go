package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Use the home function as the handler for the "/" path
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files from the "ui/static" directory.
	// note how the path is relative to the directory root
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Listening on :4000")
	// Use http.ListenAndServe to start the server on the specified port using our created mux.
	// If it returns an error then log it and exit the program.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
