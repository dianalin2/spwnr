package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"mc.honki.ng/spwnr/logging"
)

var db *sql.DB

func Init(logger logging.Logger) {
	var err error
	db, err = sql.Open("sqlite3", "database/db.sqlite3")
	if err != nil {
		logger.Log("Database failed to connect")
		return
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	CREATE TABLE IF NOT EXISTS servers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		logger.Log("Database failed to make initializing query: " + err.Error())
		return
	}
	logger.Log("Created database connection and initialized database successfully")
}

