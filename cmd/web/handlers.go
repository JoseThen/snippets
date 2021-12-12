package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/JoseThen/snippets/pkg/models"
)

// Define a handler function for the homepage that writes "Hello from Snippetbox"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
	// Use r.ParseForm() which adds any request body data for a post
	// to the r.PostForm map. Works for PUT and PATCH as well.
	// If any errors we use our ClientError helper to send a 400.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// Map for validation errors
	errors := make(map[string]string)
	// Validate Title Field
	if strings.TrimSpace(title) == "" {
		errors["title"] = "Title is required"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "Title is too long (maximum is 100 characters)"
	}

	// Validate Content Field
	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	// Validate Expires Field
	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "Invalid value for expires"
	}

	// If any errors on validation inform the user and return to the form
	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		fmt.Fprint(w, errors)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

// Add a new createSnippetForm handler, which for now returns a placeholder response.
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}
