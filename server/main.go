package main

import (
	"database/sql"
	"log"
	"server/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../addrbin.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id TEXT NOT NULL,
			password TEXT NOT NULL,
			amount INTEGER NOT NULL,
			created_at DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%S%fZ', 'now'))
		);
	`)
	if err != nil {
		log.Fatalln(err)
	}
	app.StartAPI(db)
	app.StartGenerator(db)
}
