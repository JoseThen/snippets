package main

import (
	"html/template"
	"path/filepath"

	"github.com/JoseThen/snippets/pkg/models"
)

// Define a templateData type to a pass to potential templates.
// Will add more fields eventually
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
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
		// template.Parse each page found into to a templateset (ts)
		ts, err := template.ParseFiles(page)
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
