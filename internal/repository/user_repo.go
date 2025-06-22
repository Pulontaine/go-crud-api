package repository

import (
	"api/internal/model"
	"database/sql"
)

func GetAllUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func GetUserByID(db *sql.DB, id string) (model.User, error) {
	var u model.User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func CreateUser(db *sql.DB, u model.User) error {
	return db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
}

func UpdateUser(db *sql.DB, id string, u model.User) error {
	_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
	return err
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
