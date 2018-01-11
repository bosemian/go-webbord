package model_test

import (
	"testing"

	"github.com/bosemian/go-webbord/model"
)

func TestFindAllTopic(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into topics (id, title) values (1, 'topic1')`)
	db.Exec(`insert into topics (id, title) values (2, 'topic2')`)
	db.Exec(`insert into topics (id, title) values (3, 'topic3')`)

	topics, err := model.FindAllTopic(db)
	if err != nil {
		t.Errorf(`FindAllTopic expected return list topics; but got %v`, err)
	}
	if l := len(topics); l != 3 {
		t.Fatalf("FindAllForum expected returned 3 topics; got %d", l)
	}
}

func TestCreateTopics(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	id, err := model.CreateTopic(db, &model.Topic{
		Title: "This is a topic for test",
	})
	if err != nil {
		t.Errorf(`CreateTopic expected return only one topic; but got %v`, err)
	}

	if id <= 0 {
		t.Fatalf("CreateTopic expected return positive id; got %d", id)
	}

	var cnt int
	db.QueryRow(`select count(*) from topics`).Scan(&cnt)
	if cnt != 1 {
		t.Fatalf("CreateTopic expected insert 1 topic; got %d", cnt)
	}
}

func TestUpdateTopic(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	err := model.UpdateTopic(db, &model.Topic{
		ID:    1,
		Title: "this topic has updated",
	})

	if err != nil {
		t.Errorf(`UpdateTopic expected nil; but got %v`, err)
	}
}

func TestDeletTopic(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into topics (id, title) values (1, 'topic1')`)

	err := model.DeleteTopic(db, 1)
	if err != nil {
		t.Errorf(`DeleteTopic expected return nil but got %v`, err)
	}
}
