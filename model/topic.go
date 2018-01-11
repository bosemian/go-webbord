package model

import (
	"database/sql"
	"time"
)

// Topic struct
type Topic struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindAlltopic return []*Topic
func FindAllTopic(db *sql.DB) ([]*Topic, error) {
	rows, err := db.Query(`
			select
				id, title, created_at, updated_at
			from topics
			order by created_at desc
		`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var topics []*Topic
	for rows.Next() {
		var topic Topic
		err := rows.Scan(
			&topic.ID,
			&topic.Title,
			&topic.CreatedAt,
			&topic.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		topics = append(topics, &topic)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return topics, nil
}

// CreateTopic return only one topic
func CreateTopic(db *sql.DB, t *Topic) (int, error) {
	var id int
	err := db.QueryRow(`
		insert into topics (
			title, created_at
		) values (
			$1, $2
		) returning id
	`, t.Title, time.Now()).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateTopic update topic by topic id without return utput
// if has error return err
func UpdateTopic(db *sql.DB, t *Topic) error {
	_, err := db.Exec(`
		update topics
		set title = $2, updated_at = $3
		where id = $1
	`, t.ID, t.Title, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// DeleteTopic delete topic from table
func DeleteTopic(db *sql.DB, id int) error {
	_, err := db.Exec(`
		delete from topics
		where id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}
