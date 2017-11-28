package model

import "time"
import "database/sql"

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
			title
		) values (
			$1
		) returning id
	`, forum.Title).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateForum update forum by data forum
// func UpdateForum(db *sql.DB) (*Forum, error) {

// }

// DeleteForum delate forum by ForumID
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
