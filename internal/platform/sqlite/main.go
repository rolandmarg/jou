package sqlite

import (
	"database/sql"

	// import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const schema = `
	CREATE TABLE IF NOT EXISTS journal (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(128) NOT NULL,
		created_at DATE NOT NULL,
		deleted_at DATE,
		UNIQUE(name)
	);

	CREATE TABLE IF NOT EXISTS note (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(128) NOT NULL,
		body TEXT,
		mood VARCHAR(64),
		created_at DATE NOT NULL,
		deleted_at DATE,
		j_id INTEGER NOT NULL REFERENCES journal(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS note_j_idx on note (j_id);

	CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		note_id INTEGER NOT NULL REFERENCES note(id) ON DELETE CASCADE,
		UNIQUE(note_id, name)
	);

	CREATE TABLE IF NOT EXISTS default_journal (
		name VARCHAR(64),
		j_id INTEGER NOT NULL REFERENCES journal(id) ON DELETE CASCADE,
		UNIQUE(name)
	);

	CREATE TABLE IF NOT EXISTS env (
		key VARCHAR(128) NOT NULL,
		value TEXT NOT NULL,
		UNIQUE(key)
	);
`

// Open a database connection and executes sql schema on it
func Open(name string) (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	_, err = DB.Exec(schema)
	if err != nil {
		return nil, err
	}

	return DB, nil
}

// OpenDB is a helper function that opens and populates production database
func OpenDB() (*sql.DB, error) {
	db, err := Open("jou.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
