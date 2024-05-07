package sqlstore

import "database/sql"

type DB struct {
	db *sql.DB
}

// NewDBTest creates a new DB instance for testing
func NewDBTest(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

//
//func NewDB(dsn string) *DB {
//	db, err := sql.Open("postgres", dsn)
//	if err != nil {
//		panic(err)
//	}
//
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	return &DB{
//		db: db,
//	}
//}
