package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func NewPostgresStore() (*sql.DB, error) {

	db, err := sql.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
