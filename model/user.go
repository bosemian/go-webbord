package model

import (
	"database/sql"
	"time"

	"github.com/bosemian/go-webbord/api"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Find User by username return username just one Row
func FindUserByUsername(db *sql.DB, username string) (*User, error) {
	var user User
	err := db.QueryRow(`
			select
				id, username, password, created_at, updated_at
			from users
			where username = $1
		`, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, api.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser insert user to db return user just only one row
func CreateUser(db *sql.DB, u *User) (int, error) {
	var id int
	err := db.QueryRow(`
		insert into users (
			id, username, password
		) values (
			$1, $2, $3
		) returning id
	`, 1, u.Username, u.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
