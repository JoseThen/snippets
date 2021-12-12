package main

import (
	"html/template"
	"net/url"
	"path/filepath"
	"time"

	"github.com/JoseThen/snippets/pkg/models"
)

// Define a templateData type to a pass to potential templates.
// Will add more fields eventually
type templateData struct {
	CurrentYear int
	FormData    url.Values
	FormErrors  map[string]string
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Init a map as a cache
	cache := map[string]*template.Template{}

	// Use `filepath.Glob` to get a slice of anything with the .page.tmpl extension
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// For each file found ...
	for _, page := range pages {
		// Save file name to variable
		name := filepath.Base(page)
		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Parse any layout files to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		// Parse any partial files to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		// Add the template set to the cache under the name of the file
		cache[name] = ts
	}

	return cache, nil
}
