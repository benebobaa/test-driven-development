package sqlstore

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB *sql.DB
)

var (
	postgresDNS = "postgres://root:root@localhost:5432/tdd?sslmode=disable"
)

func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", postgresDNS)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	testDB = db

	// Truncate the users table before running the tests
	_, err = testDB.Exec("TRUNCATE users")
	if err != nil {
		log.Fatal(err)

	}
	code := m.Run()

	// Close the database connection after running the tests
	err = testDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
