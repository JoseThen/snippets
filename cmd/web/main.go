package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/JoseThen/snippets/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

// Struct to hold application wide dependencies/configuration
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	// Define a flag for the port the application runs on
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Flag for the MySQL DSN string
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "MySQL data source name")
	// Parse the command line flag. Need to call before you use the variable.
	flag.Parse()

	// Make an custom logger for info prefixed with info, and adding some more data like date and time
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)

	// Create an error logger, but write to stderr and use `Lshortfile` to include file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// Close the db before the main function exits
	defer db.Close()

	// Init an new instance of the application containing dependencies
	// Initialize a mysql.SnippetModel instance and add it to the application
	// dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
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
	// `err =` instead of `err :=` because its declared already above
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
