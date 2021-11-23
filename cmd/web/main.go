package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define a flag for the port the application runs on
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the command line flag. Need to call before you use the variable.
	flag.Parse()

	// Make an custom logger for info prefixed with info, and adding some more data like date and time
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)

	// Create an error logger, but write to stderr and use `Lshortfile` to include file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

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

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Value returned by `flag.String()` is a pointer, not the actual value so we devalue that with '*'
	infoLog.Printf("Listening on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	// If it returns an error then log it and exit the program.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
