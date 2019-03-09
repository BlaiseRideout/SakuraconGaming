package gameroom

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./gameroom.sqlite3?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
}
