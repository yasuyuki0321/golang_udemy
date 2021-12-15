package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ = sql.Open("sqlite3", "example.sql")
	defer Db.Close()

	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "hanako", 19)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cmd := "SELECT * FROM persons WHERE age = ?"
	row := Db.QueryRow(cmd, 40)

	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err == sql.ErrNoRows {
		log.Println("No row")
	} else {
		log.Println(err)
	}
	fmt.Println(p.Name, p.Age)
}
