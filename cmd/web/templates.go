package main

import "github.com/JoseThen/snippets/pkg/models"

// Define a templateData type to a pass to potential templates.
// Will add more fields eventually
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
