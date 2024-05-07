package sqlstore

import (
	"github.com/lib/pq"
	"test-driven-development/pkg/common"
	"test-driven-development/pkg/domain"
)

var (
	sqlCreateUser = `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
)

func (d *DB) CreateUser(user *domain.User) (*domain.User, error) {

	// Query create user and then scan the id
	err := d.db.QueryRow(sqlCreateUser, user.Username, user.Password).Scan(&user.ID)

	// Handle unique pqError code with 23505 is unique constraint violation
	pgErr, ok := err.(*pq.Error)
	if ok && pgErr.Code == "23505" {
		// Handle unique constraint violation error here
		// For example, you can return a custom error indicating the username is already taken
		return nil, common.ErrDuplicateUsername
	}

	// Return error if any
	if err != nil {
		return nil, err
	}

	// Return the user
	return user, nil
}
