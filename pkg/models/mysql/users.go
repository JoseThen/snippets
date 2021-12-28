package mysql

import (
	"database/sql"

	"github.com/JoseThen/snippets/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Function to create a user into the DB. This should return a nil error
// if all went well.
func (userModel *UserModel) Insert(name, email, password string) error {
	return nil
}

// Function to login the user and check against DB. Will return the User ID if
// if passes.
func (userModel *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Fetch details for the specific  user based on their ID.
func (userModel *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
