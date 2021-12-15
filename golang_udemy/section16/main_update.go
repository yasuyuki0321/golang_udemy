package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	Db, _ = sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	cmd := "UPDATE persons SET age = ? WHERE name = ?"
	_, err := Db.Exec(cmd, 40, "tarou")
	if err != nil {
		log.Fatal(err)
	}
}
