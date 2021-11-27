package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JoseThen/snippets/pkg/models"
)

// Define a handler function for the homepage that writes "Hello from Snippetbox"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because the "/" path is a catchall and will route here, we need to check if the request is for the homepage
	if r.URL.Path != "/" {
		// If it isn't the homepage, then return a 404 not found error
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: snippets})
}

// showSnippet handler function
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Fetch the ID from the URL query param and if its less than 1 or not a number, return a 400 Bad Request error
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: snippet})
}

// createSnippet handler function
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != "POST" {
		// If not a post, respond with which methods are allowed
		w.Header().Set("Allow", http.MethodPost)
		// Return a 405 Method Not Allowed error
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Dummy Data
	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
