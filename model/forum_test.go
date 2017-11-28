package model_test

import (
	"testing"

	"github.com/bosemian/go-webbord/model"
)

func TestFindAllForum(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into forums (id, title) values (1, 'title1')`)
	db.Exec(`insert into forums (id, title) values (2, 'title2')`)
	db.Exec(`insert into forums (id, title) values (3, 'title3')`)

	forums, err := model.FindAllForum(db)
	if err != nil {
		t.Fatalf("FindAllForum expected returned nil error; got %v", err)
	}
	if l := len(forums); l != 3 {
		t.Fatalf("FindAllForum expected returned 3 forums; got %d", l)
	}
}

func TestCreateForum(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	id, err := model.CreateForum(db, &model.Forum{
		Title: "title test",
	})

	if err != nil {
		t.Fatalf("CreateForum expected return nil error: got %v", err)
	}

	if id < 0 {
		t.Fatalf("CreateForum expected return positive id; got %d", id)
	}

	var cnt int
	db.QueryRow(`select count(*) from forums`).Scan(&cnt)
	if cnt != 1 {
		t.Fatalf("CreateForum expected insert 1 forum; got %d", cnt)
	}
}

func TestDeleteForum(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	id, err := model.CreateForum(db, &model.Forum{
		Title: "title test",
	})
	err = model.DeleteForum(db, id)
	if err != nil {
		t.Fatalf("DeletedForum expected return 1 got; %v", err)
		return
	}
}
