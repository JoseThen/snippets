package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JoseThen/snippets/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

// Struct to hold application wide dependencies/configuration
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
	users         *mysql.UserModel
}

func main() {
	// Define a flag for the port the application runs on
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Flag for the MySQL DSN string
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "MySQL data source name")
	// New flag for sessions secret (random key) for encryption, should be 32 bytes long
	secret := flag.String("secret", "s6Ndh+3rdfKDMs*+9PKILOPa7qGWhTzbpa@ge", "Secret Key for Session")
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

	// Inite a new template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Use sessions.New() to init a sessions manager
	// configure so it expires after 12 hours
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true // Set secure flaf for sessions cookies

	// Init an new instance of the application containing dependencies
	// Initialize a mysql.SnippetModel instance and add it to the application
	// dependencies.
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		// Add users to application struct
		users: &mysql.UserModel{DB: db},
	}

	// Init a tls.Config struct to hold non default tls settings
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	// Set the server tls config for the http.Server struct
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Value returned by `flag.String()` is a pointer, not the actual value so we devalue that with '*'
	infoLog.Printf("Listening on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	// If it returns an error then log it and exit the program.
	// `err =` instead of `err :=` because its declared already above
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
