package model

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindAllComment find every row in comment table
// return all comment
func FindAllComment(db *sql.DB) ([]*Comment, error) {
	rows, err := db.Query(`
	 select id, comment, created_at, updated_at
	 from comments
	 order by created_at desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []*Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Comment,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		comments = append(comments, &comment)
		if err != nil {
			return nil, err
		}
	}
	return comments, nil
}

// CreateComment insert comment and return id of inserted
func CreateComment(db *sql.DB, c *Comment) (int, error) {
	var id int
	err := db.QueryRow(`
		insert comments
		into (comment, created_at)
		values ($1, $2)
	`, c.Comment, c.CreatedAt).Scan(&id)
	if err != nil {
		return 0, nil
	}
	return id, nil
}

// UpdateComment change comment and time by id
func UpdateComment(db *sql.DB, c *Comment) error {
	row, err := db.Exec(`
		update comments
		set comment = $2, updated_at = $3
		where id = $1
	`, c.ID, c.Comment, c.UpdatedAt)
	if err != nil {
		return err
	}
	_, err = row.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment delete from comment table
func DeleteComment(db *sql.DB, id int) error {
	row, err := db.Exec(`
		delete from comments
		where id = $1
	`, id)
	if err != nil {
		return err
	}
	_, err = row.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
