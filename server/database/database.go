package database

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"mc.honki.ng/spwnr/logging"
)

var db *sql.DB
var logger logging.Logger

func Init(loggerArg logging.Logger) {
	logger = loggerArg

	db, err := sql.Open("sqlite3", "database/db.sqlite3")
	if err != nil {
		logger.Log("Database failed to connect")
		return
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS servers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		status TEXT NOT NULL,

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

func CreateUser(email string, username string, password string) (int64, error) {
	query := `
	INSERT INTO users(email, username, password)
	VALUES(?, ?, ?)
	`
	result, err := db.Exec(query, email, username, password)
	if err != nil {
		logger.Log("Failed to create user")
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Log("Failed to get id while creating user")
		return -1, err
	}
	return id, nil
}

func CreateServer(user_id int64, name string) (int64, error) {
	query := `
	INSERT INTO servers(user_id, name)
	VALUES(?, ?)
	`
	result, err := db.Exec(query, user_id, name)
	if err != nil {
		logger.Log("Failed to create server")
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Log("Failed to get id while creating server")
		return -1, err
	}
	return id, nil
}

func SetServerStatus(id int64, newStatus string) error {
	query := `
	UPDATE servers SET status = ?
	WHERE id = ?
	`
	_, err := db.Exec(query, newStatus, id)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to update server (id=%d) with new status '%s'", id, newStatus))
		return err
	}
	return nil
}
