package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a flag for the port the application runs on
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the command line flag. Need to call before you use the variable.
	flag.Parse()
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

	// Value returned by `flag.String()` is a pointer, not the actual value so we devalue that with '*'
	log.Printf("Listening on %s", *addr)
	// Use http.ListenAndServe to start the server on the specified port using our created mux.
	// If it returns an error then log it and exit the program.
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
