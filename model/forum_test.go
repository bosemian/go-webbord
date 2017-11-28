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
