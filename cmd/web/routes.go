package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// Function returns an httphandler so that we can use it with our middleware
func (app *application) routes() http.Handler {
	// Create middleware chain containing the standard middleware for every reqyest
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// Create a new ServeMux
	mux := pat.New()
	// Use the home function as the handler for the "/" path
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	// Create a file server which serves files from the "ui/static" directory.
	// note how the path is relative to the directory root
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Return the ServeMux
	// Return the standardMiddlware variable/chain with the mux as the final handler
	return standardMiddleware.Then(mux)
}
