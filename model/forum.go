package model

import (
	"database/sql"
	"time"
)

type Forum struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindAllForum find all forum
func FindAllForum(db *sql.DB) ([]*Forum, error) {
	rows, err := db.Query(`
			select
				id, title, created_at, updated_at
			from forums
			order by created_at desc
		`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var forums []*Forum
	for rows.Next() {
		var forum Forum
		err := rows.Scan(
			&forum.ID,
			&forum.Title,
			&forum.CreatedAt,
			&forum.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		forums = append(forums, &forum)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return forums, nil
}

// CreateForum create a new Forum
func CreateForum(db *sql.DB, forum *Forum) (int, error) {
	var id int
	err := db.QueryRow(`
		insert into forums (
			title, created_at
		) values (
			$1, $2
		) returning id
	`, forum.Title, time.Now()).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateForum update forum by data forum
func UpdateForum(db *sql.DB, f *Forum) error {
	_, err := db.Exec(`
		update forums
		set title = $2, updated_at = $3
		where id = $1
	`, f.ID, f.Title, f.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteForum delete forum by ForumID
func DeleteForum(db *sql.DB, forumID int) error {
	_, err := db.Exec(`
		delete from forums
		where id = $1
	`, forumID)
	if err != nil {
		return err
	}
	return nil
}

// FindForumByID return a just only forum
func FindForumByID(db *sql.DB, id int) (*Forum, error) {
	var forum Forum
	err := db.QueryRow(`
		select
			id, title, created_at, updated_at
		from forums
		where id = $1
	`, id).Scan(&forum.ID, &forum.Title, &forum.CreatedAt, &forum.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &forum, nil
}
