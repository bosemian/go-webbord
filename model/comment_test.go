package model_test

import (
	"testing"
	"time"

	"github.com/bosemian/go-webbord/model"
)

func TestFindAllComment(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into comments (id, comment) values (1, 'comment1')`)
	db.Exec(`insert into comments (id, comment) values (2, 'comment2')`)
	db.Exec(`insert into comments (id, comment) values (3, 'comment3')`)

	cmnt, err := model.FindAllComment(db)
	if err != nil {
		t.Errorf(`FindAllComment expected list comment; but got %v`, err)
	}
	if len(cmnt) != 3 {
		t.Errorf(`FindAllComment expected just 3 comment; but got %v`, err)
	}
}

func TestCreateComment(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	id, err := model.CreateComment(db, &model.Comment{
		ID:        1,
		Comment:   "This is a comment",
		CreatedAt: time.Now(),
	})
	if err != nil {
		t.Errorf(`CreateComment expected nil; but got %v`, err)
	}
	if id < 0 {
		t.Errorf(`CreateComment expected id value is not 0`)
	}
}

func TestUpdateComment(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into comments (id, comment) values (1, 'comment1')`)

	err := model.UpdateComment(db, &model.Comment{
		ID:        1,
		Comment:   "This comment has updated",
		UpdatedAt: time.Now(),
	})

	if err != nil {
		t.Errorf(`UpdateComment expected return nil; but got %v`, err)
	}
}

func TestDeleteComment(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into comments (id, comment) values (1, 'comment1')`)
	err := model.DeleteComment(db, 2)
	if err != nil {
		t.Errorf(`DeleteComment expected return nil; but got %v`, err)
	}
}
