package sqlstore

import (
	"github.com/stretchr/testify/assert"
	"test-driven-development/pkg/common"
	"test-driven-development/pkg/domain"
	"testing"
)

func TestCreateUser(t *testing.T) {

	sqlStore := NewDBTest(testDB)

	user := &domain.User{
		Username: "bene",
		Password: "password",
	}

	createdUser, err := sqlStore.CreateUser(user)

	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Password, createdUser.Password)
}

func TestCreateUserDuplicateUsername(t *testing.T) {
	sqlStore := NewDBTest(testDB)

	user := &domain.User{
		Username: "test",
		Password: "password",
	}

	// Create a user
	_, err := sqlStore.CreateUser(user)
	assert.Nil(t, err)

	// Create another user with the same username
	createdUserTwo, err := sqlStore.CreateUser(user)

	assert.NotNil(t, err)
	// Check if the error is ErrDuplicateUsername
	assert.Equal(t, err, common.ErrDuplicateUsername)
	assert.Nil(t, createdUserTwo)
}

func TestThrowErrorInvalidSQL(t *testing.T) {
	oldSql := sqlCreateUser

	sqlStore := NewDBTest(testDB)

	user := &domain.User{
		Username: "bene",
		Password: "password",
	}

	// Inject invalid SQL
	sqlCreateUser = `invalid sql query`

	createdUser, err := sqlStore.CreateUser(user)
	assert.Nil(t, createdUser)
	assert.NotNil(t, err)

	// Reset the sqlCreateUser
	sqlCreateUser = oldSql
}
