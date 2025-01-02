package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Couldn't connect database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()

}

func createTable() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL 
	)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL, 
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER,
	event_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(event_id) REFERENCES events(id)
	)
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

}
