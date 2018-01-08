package model_test

import (
	"testing"

	"github.com/bosemian/go-webbord/model"
)

func TestFindUserByUsername(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into users (id, username, password) values (1, 'user1', 'password')`)

	_, err := model.FindUserByUsername(db, "user1")
	if err != nil {
		t.Fatalf("FindUserByUsername expected returned nil error; got %v", err)
	}
	var cnt int
	db.QueryRow(`select count(*) from users`).Scan(&cnt)
	if cnt != 1 {
		t.Fatalf("FindUserByUsername expected retrun 1 users; got %d", cnt)
	}
}

func TestCreateUser(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	_, err := model.CreateUser(db, &model.User{
		Username: "user1",
		Password: "password1",
	})
	if err != nil {
		t.Fatalf(`TestCreateUser expected return nil; but got %v`, err)
	}
	var cnt int
	db.QueryRow(`select count(*) from users`).Scan(&cnt)
	if cnt != 1 {
		t.Fatalf("TestCreateUser expected retrun 1 users; got %d", cnt)
	}
}
