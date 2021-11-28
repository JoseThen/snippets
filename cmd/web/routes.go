package main

import "net/http"

// Function returns an httphandler so that we can use it with our middleware
func (app *application) routes() http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Use the home function as the handler for the "/" path
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files from the "ui/static" directory.
	// note how the path is relative to the directory root
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Return the ServeMux
	// Chain middlewares
	return app.logRequest(secureHeaders(mux))
}
