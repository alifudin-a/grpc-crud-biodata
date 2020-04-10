package app

import (
	"github.com/jmoiron/sqlx"
)

// App : will be used later
type App struct {
	Dbx *sqlx.DB
}

// OpenDB : open the database connection
func OpenDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
