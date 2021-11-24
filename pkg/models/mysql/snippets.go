package mysql

import (
	"database/sql"

	"github.com/JoseThen/snippets/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Create a new snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get a snippet by its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Fetch latest 10 snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
