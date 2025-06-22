package db

import "database/sql"

func Connect(url string) (*sql.DB, error) {
	return sql.Open("postgres", url)
}

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT
	)`
	_, err := db.Exec(query)
	return err
}
