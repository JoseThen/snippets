package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Struct to hold application wide dependencies/configuration
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a flag for the port the application runs on
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the command line flag. Need to call before you use the variable.
	flag.Parse()

	// Make an custom logger for info prefixed with info, and adding some more data like date and time
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)

	// Create an error logger, but write to stderr and use `Lshortfile` to include file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	// Init an new instance of the application containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Value returned by `flag.String()` is a pointer, not the actual value so we devalue that with '*'
	infoLog.Printf("Listening on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	// If it returns an error then log it and exit the program.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
