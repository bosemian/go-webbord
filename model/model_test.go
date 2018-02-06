package model_test

import (
	"database/sql"
	"io/ioutil"
	"testing"

	_ "github.com/lib/pq"
)

// const PATH_DB = "postgres://postgres@localhost:5432?sslmode=disable"

const PATH_DB_DOCKER = "postgres://postgres:password@192.168.99.100:5432?sslmode=disable"

func prepareDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", PATH_DB)
	if err != nil {
		t.Fatal(err)
	}

	// gowebbord_test
	_, err = db.Exec("drop database if exists gowebbord_test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`create database gowebbord_test`)
	if err != nil {
		t.Fatal(err)
	}

	db.Close()

	db, err = sql.Open("postgres", "postgres://postgres:password@localhost:5432/gowebbord_test?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	schema, err := ioutil.ReadFile("../model.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		t.Fatal(err)
	}

	return db
}
