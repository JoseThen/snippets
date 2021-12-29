package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/JoseThen/snippets/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

// Function to create a user into the DB. This should return a nil error
// if all went well.
func (userModel *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = userModel.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		// If this returns an error, we use the errors.As() function to check
		// whether the error has the type *mysql.MySQLError. If it does, the
		// error will be assigned to the mySQLError variable. We can then check
		// whether or not the error relates to our users_uc_email key by
		// checking the contents of the message string. If it does, we return
		// an ErrDuplicateEmail error.
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
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
